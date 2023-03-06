package signals

import (
	"os"
	"os/signal"
)

// New returns wrapper of signal.Notify
func New() <-chan struct{} {
	out := make(chan struct{})

	go func() {
		si := make(chan os.Signal, 1)
		signal.Notify(si, os.Interrupt, os.Kill)
		<-si

		close(out)
	}()

	return out
}
