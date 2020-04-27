package client

import (
	"github.com/9299381/wego/clients"
	"testing"
)

func TestClientService(t *testing.T) {
	param := make(map[string]interface{})
	resp := clients.Service("demo").
		Api("demo.post").
		Params(param).
		Run()

	t.Log(resp)
}
