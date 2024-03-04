package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"go-project/config"
)

type Server struct {
	cfg    *config.Config
	echo   *echo.Echo
	logger *logrus.Logger
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg, echo: echo.New(), logger: logrus.New()}
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:         ":" + s.cfg.Server.Port,
		WriteTimeout: s.cfg.Server.Timeout.Write * time.Second,
		ReadTimeout:  s.cfg.Server.Timeout.Read * time.Second,
	}

	go func() {
		s.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalln("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	// Set up a channel to listen to interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTERM)

	// Block on this channel listening for those previously defined syscalls assign
	// to variable, so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		s.cfg.Server.Timeout.Server,
	)
	defer cancel()

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	s.logger.Println("Server is shutting down due to", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		s.logger.Fatalln("Server was unable to gracefully shutdown due to err:", err)
	}

	s.logger.Fatalln("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
