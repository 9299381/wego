package queues

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/configs"
)

func Fire(name string, router string, params map[string]interface{}) error {
	conn := clients.Redis()
	defer conn.Close()
	job := &Job{
		Queue: name,
		Payload: Payload{
			Route:  router,
			Params: params,
		},
	}
	prefix := configs.EnvString("queue.prefix", "wego")
	return Enqueue(conn, job, prefix)

}
