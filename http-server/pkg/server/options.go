package server

import (
	"net/http"

	"github.com/MedvedewEM/pow/pkg/signals"
)

// Option is a way to set up server.
// There are required internal options and optional public options
type Option func(*srv)

func gracefulShutdownOption() Option {
	return func(s *srv) {
		s.stop = signals.New()
	}
}

func httpServerOption() Option {
	return func(s *srv) {
		s.httpServer = http.Server{}
	}
}

func errorOutOption() Option {
	return func(s *srv) {
		s.errOut = make(chan error)
	}
}

// HttpServerOption can be use in case of http-server is required
func HttpServerOption(addr string, handler http.Handler) Option {
	return func(s *srv) {
		s.httpServer.Addr = addr
		s.httpServer.Handler = handler
	}
}
