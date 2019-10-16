package cronjobs

import (
	"github.com/9299381/wego/servers/commons"
	"github.com/robfig/cron/v3"
	"time"
)

type Server struct {
	Serv *cron.Cron
}

func NewServer() *Server {
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	ss := &Server{
		Serv: cron.New(cron.WithSeconds(), cron.WithLocation(nyc)),
	}
	return ss
}

func (s *Server) Register(spec string, job *commons.CommHandler) {
	_, _ = s.Serv.AddJob(spec, job)
}

func (s *Server) Serve() error {
	s.Serv.Start()
	select {}
}
func (s *Server) Close() {

}
