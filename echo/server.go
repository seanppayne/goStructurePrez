package echo

import (
	"context"
	"example.com/demo"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type echoServer struct {
	server *echo.Echo
	Options
}

type Options struct {
	EchoConfig *demo.EchoConfig

	UserRepo demo.UserRepository
}

// NewServer creates a new instance of the echoServer.
func NewServer(opts Options) demo.Server {
	return &echoServer{
		Options: opts,
	}
}

func (s *echoServer) Run(ctx context.Context, wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()
		<-ctx.Done()
		s.server.Close()
	}()

	if err := s.server.Start(":8080"); err != http.ErrServerClosed {
		return err
	}

	return nil
}

// RegisterHandlers registers the handlers for the echoServer.
func (s *echoServer) RegisterHandlers() {
	s.server = echo.New()
	s.RegisterUserHandlers()
}
