package configs

import "strconv"

type EventConfig struct {
	Concurrency int `json:"concurrency"`
	After       int `json:"after"`
}

func (s *EventConfig) Load() *EventConfig {
	concurrency, _ := strconv.Atoi(Env("EVENT_CONCURRENCY", "1"))
	after, _ := strconv.Atoi(Env("EVENT_AFTER", "1"))

	config := &EventConfig{
		Concurrency: concurrency,
		After:       after,
	}
	return config
}
