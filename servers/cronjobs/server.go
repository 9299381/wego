package cronjobs

import (
	"github.com/9299381/wego/servers/commons"
	"github.com/robfig/cron/v3"
	"time"
)

type Server struct {
	Server *cron.Cron
}

func NewServer() *Server {
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	ss := &Server{
		Server: cron.New(cron.WithSeconds(), cron.WithLocation(nyc)),
	}
	return ss
}

func (it *Server) Register(spec string, job *commons.CommHandler) {
	_, _ = it.Server.AddJob(spec, job)
}

func (it *Server) Serve() error {
	it.Server.Start()
	select {}
}
