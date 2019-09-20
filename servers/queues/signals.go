package queues

import (
	"os"
	"os/signal"
	"syscall"
)

func signals() <-chan bool {
	quit := make(chan bool)

	go func() {
		signals := make(chan os.Signal)
		defer close(signals)

		signal.Notify(signals, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt)
		defer signalStop(signals)

		<-signals
		quit <- true
	}()

	return quit
}
