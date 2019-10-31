package configs

type EventConfig struct {
	Concurrency int `json:"concurrency"`
	After       int `json:"after"`
}

func LoadEventConfig() *EventConfig {

	config := &EventConfig{
		Concurrency: EnvInt("event.concurrency", 1),
		After:       EnvInt("event.after", 1),
	}
	return config
}
