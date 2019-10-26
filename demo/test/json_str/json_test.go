package json_str

import (
	"encoding/json"
	"fmt"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/idwork"
	"testing"
)

type DemoStruct struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func TestDemoStruct(t *testing.T) {
	var records []*DemoStruct
	records = append(records, &DemoStruct{
		Id:   idwork.ID(),
		Name: "one",
	})
	records = append(records, &DemoStruct{
		Id:   idwork.ID(),
		Name: "two",
	})
	args, _ := json.Marshal(records)
	fmt.Println(string(args))
	params := map[string]interface{}{
		"args": string(args),
	}

	payload := &contracts.Payload{
		Route:  "route",
		Params: params,
	}
	jb, _ := json.Marshal(payload)

	p := &contracts.Payload{}
	_ = json.Unmarshal(jb, p)
	fmt.Println(p)

	var st []*DemoStruct
	_ = json.Unmarshal([]byte(p.Params["args"].(string)), &st)
	for _, v := range st {
		fmt.Println(v)
	}

}
