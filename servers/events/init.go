package events

import "github.com/go-kit/kit/endpoint"

var Event *eventStack
var Handlers map[string]endpoint.Endpoint

func init() {
	Handlers = map[string]endpoint.Endpoint{}
	Event = newEventStack()
}
