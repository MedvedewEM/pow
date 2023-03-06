package server

import (
	"context"
	"net/http"
)

// Interface defines methods to control server.
type Interface interface {
	// Run method runs server and returns readonly channel for error handling.
	Run() <-chan error

	// Wait method waits while server working or until server will be shutdowned.
	Wait()
}

// New returns new srv app to run.
func New(opts ...Option) *srv {
	s := &srv{}

	opts = addRequiredOptions(opts)

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type srv struct {
	httpServer http.Server

	stop   <-chan struct{}
	errOut chan error
}

func (s *srv) Run() <-chan error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			s.errOut <- err
		}
	}()

	return s.errOut
}

func (s *srv) Wait() {
	<-s.stop

	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		s.errOut <- err
	}
}

func addRequiredOptions(opts []Option) []Option {
	reqs := []Option{
		gracefulShutdownOption(),
		httpServerOption(),
		errorOutOption(),
	}

	return append(reqs, opts...)
}
