package events

import (
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
	"runtime"
	"sync"
)

/**
通过channel方式传递event,而不是通过共享内存传递
*/

var Handlers map[string]endpoint.Endpoint
var eventPool sync.Pool
var eventChan chan *contracts.Payload

func init() {
	Handlers = make(map[string]endpoint.Endpoint)
	eventChan = make(chan *contracts.Payload, runtime.NumCPU())
}
