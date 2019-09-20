// +build go1.1

package queues

import (
	"os"
	"os/signal"
)
func signalStop(c chan<- os.Signal) {
	signal.Stop(c)
}
