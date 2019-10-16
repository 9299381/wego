package queues

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/commons"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Options struct {
	Prefix      string
	Listen      []string
	Interval    time.Duration
	Concurrency int
	UseNumber   bool
}

type Server struct {
	opts      Options
	handlers  map[string]*commons.CommHandler
	ctx       context.Context
	RedisPool *redis.Pool
	Logger    contracts.ILogger
}

func NewServer(opts *Options) *Server {
	//初始化,logger,redis池
	s := &Server{
		opts:     *opts,
		ctx:      context.Background(),
		handlers: make(map[string]*commons.CommHandler),
	}
	return s
}

func (s *Server) Register(name string, handler *commons.CommHandler) {
	s.handlers[name] = handler

}

func (s *Server) Serve() error {
	errChan := make(chan error)
	quit := signals()
	jobs := s.poll(quit, errChan)
	for id := 0; id < s.opts.Concurrency; id++ {
		s.work(id, jobs, errChan)
	}
	return <-errChan
}

func (s *Server) work(id int, jobs <-chan *Job, errChan chan error) {
	go func() {
		for job := range jobs {
			if handler, ok := s.handlers[job.Payload.Route]; ok {
				ctx := context.Background()
				ctx = context.WithValue(ctx, "Queue", job.Queue)
				request := job.Payload.Params
				response, err := handler.Handle(ctx, request)
				if err != nil {
					errChan <- err
					return
				}
				msg := fmt.Sprintf(
					"Concurrency_Id:%d ,Job Response:%v",
					id,
					response)
				s.Logger.Debug(msg)
			} else {
				errorLog := fmt.Sprintf(
					"No worker for %s in queue %s with args %v",
					job.Payload.Route,
					job.Queue,
					job.Payload.Params)
				s.Logger.Error(errorLog)
				errChan <- errors.New(errorLog)
				return
			}
		}
	}()
}

func (s *Server) poll(quit <-chan bool, errChan chan error) <-chan *Job {
	jobs := make(chan *Job)
	go func() {
		conn := s.RedisPool.Get()
		defer conn.Close()
		for {
			select {
			default:
				job, err := s.getJob(conn)
				if err != nil {
					errorLog := fmt.Sprintf(
						"Error on %v getting job from: %v",
						s.opts.Listen,
						err)
					s.Logger.Error(errorLog)
					errChan <- errors.New(errorLog)
					return
				}
				if job != nil {
					select {
					case jobs <- job:
					case <-quit:
						buf, err := json.Marshal(job.Payload)
						if err != nil {
							errorLog := fmt.Sprintf(
								"Error requeueing %v: %v",
								job, err)
							s.Logger.Error(errorLog)
							errChan <- errors.New(errorLog)
							return
						}
						arg := fmt.Sprintf("%s_queue:%s", s.opts.Prefix, job.Queue)
						_ = conn.Send("LPUSH", arg, buf)
						_ = conn.Flush()
						return
					}
				} else {
					s.Logger.Debugf("Sleeping for %v", s.opts.Interval)
					s.Logger.Debugf("Waiting for %v", s.opts.Listen)
					timeout := time.After(s.opts.Interval)
					select {
					case <-quit:
						return
					case <-timeout:
					}
				}
			}
		}
	}()

	return jobs
}

func (s *Server) getJob(conn redis.Conn) (*Job, error) {
	for _, queue := range s.opts.Listen {
		s.Logger.Debugf("Checking %s", queue)
		arg := fmt.Sprintf("%s_queue:%s", s.opts.Prefix, queue)
		reply, err := conn.Do("LPOP", arg)
		if err != nil {
			return nil, err
		}
		if reply != nil {
			s.Logger.Debugf("Found job on %s", queue)
			job := &Job{Queue: queue}
			decoder := json.NewDecoder(bytes.NewReader(reply.([]byte)))
			if s.opts.UseNumber {
				decoder.UseNumber()
			}
			if err := decoder.Decode(&job.Payload); err != nil {
				return nil, err
			}
			return job, nil
		}
	}
	return nil, nil
}

func (s *Server) Close() {

}
