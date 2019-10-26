package queue

import (
	"encoding/json"
	"github.com/9299381/wego/demo/src/dto"
	"github.com/9299381/wego/servers/queues"
	"github.com/9299381/wego/tools/idwork"
	"testing"
)

func TestQueue(t *testing.T) {
	var orders []*dto.Order
	orders = append(orders, &dto.Order{
		Id:   idwork.ID(),
		Name: "one",
	})
	orders = append(orders, &dto.Order{
		Id:   idwork.ID(),
		Name: "two",
	})
	orderJson, _ := json.Marshal(orders)

	_ = queues.Fire(
		"demo1",
		"queue2",
		map[string]interface{}{
			"task_id":    idwork.ID(),
			"order_json": string(orderJson),
		},
	)

}
