package app

import (
	"log"

	"github.com/MedvedewEM/pow/internal/internal/challenge"
	"github.com/MedvedewEM/pow/internal/internal/wisdom"
	"github.com/MedvedewEM/pow/pkg/server"
)

func New() *App {
	config := NewConfig()

	challengeService := challenge.NewChallenge()
	wisdomService := wisdom.NewWisdom()

	auth := NewAuthMiddleware(challengeService)

	handler := NewHandler(auth.Middleware, challengeService, wisdomService)

	srv := server.New(
		server.HttpServerOption(config.Addr, handler),
	)

	return &App{
		srv: srv,
	}
}

type App struct {
	srv server.Interface
}

func (p *App) Run() {
	errOut := p.srv.Run()

	go func() {
		err := <-errOut
		if err != nil {
			log.Fatalln("Error from http-server:", err)
		}
	}()

	p.srv.Wait()
}
