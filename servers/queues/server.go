package queues

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/9299381/wego/servers/commons"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
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
	Logger    *logrus.Logger
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

func (it *Server) Register(name string, handler *commons.CommHandler) {
	it.handlers[name] = handler

}

func (it *Server) Serve() error {
	errChan := make(chan error)
	quit := signals()
	jobs := it.poll(quit, errChan)
	for id := 0; id < it.opts.Concurrency; id++ {
		it.work(id, jobs, errChan)
	}
	return <-errChan
}

func (it *Server) work(id int, jobs <-chan *Job, errChan chan error) {
	go func() {
		for job := range jobs {
			if handler, ok := it.handlers[job.Payload.Route]; ok {
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
				it.Logger.Debug(msg)
			} else {
				errorLog := fmt.Sprintf(
					"No worker for %s in queue %s with args %v",
					job.Payload.Route,
					job.Queue,
					job.Payload.Params)
				it.Logger.Error(errorLog)
				errChan <- errors.New(errorLog)
				return
			}
		}
	}()
}

func (it *Server) poll(quit <-chan bool, errChan chan error) <-chan *Job {
	jobs := make(chan *Job)
	go func() {
		conn := it.RedisPool.Get()
		defer conn.Close()
		for {
			select {
			default:
				job, err := it.getJob(conn)
				if err != nil {
					errorLog := fmt.Sprintf(
						"Error on %v getting job from: %v",
						it.opts.Listen,
						err)
					it.Logger.Error(errorLog)
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
							it.Logger.Error(errorLog)
							errChan <- errors.New(errorLog)
							return
						}
						arg := fmt.Sprintf("%s_queue:%s", it.opts.Prefix, job.Queue)
						_ = conn.Send("LPUSH", arg, buf)
						_ = conn.Flush()
						return
					}
				} else {
					it.Logger.Debugf("Sleeping for %v", it.opts.Interval)
					it.Logger.Debugf("Waiting for %v", it.opts.Listen)
					timeout := time.After(it.opts.Interval)
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

func (it *Server) getJob(conn redis.Conn) (*Job, error) {
	for _, queue := range it.opts.Listen {
		it.Logger.Debugf("Checking %s", queue)
		arg := fmt.Sprintf("%s_queue:%s", it.opts.Prefix, queue)
		reply, err := conn.Do("LPOP", arg)
		if err != nil {
			return nil, err
		}
		if reply != nil {
			it.Logger.Debugf("Found job on %s", queue)
			job := &Job{Queue: queue}
			decoder := json.NewDecoder(bytes.NewReader(reply.([]byte)))
			if it.opts.UseNumber {
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
