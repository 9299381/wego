package queues

type Payload struct {
	Route string `json:"route"`
	Params  map[string]interface{} `json:"params"`
}
