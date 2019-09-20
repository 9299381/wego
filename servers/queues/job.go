package queues

type Job struct {
	Queue   string
	Payload Payload
}
