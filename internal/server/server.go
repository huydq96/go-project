package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-project/internal/alert/delivery/http/v1"
	"go-project/internal/repository/mongo"
	"go-project/internal/usecase"
	"go-project/pkg/interfaces/database/mongodb"

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
	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		s.cfg.Server.Timeout.Server,
	)
	defer cancel()

	// mongodb client
	mongoAlertClient, err := mongodb.NewClient(s.cfg.Adapter.Mongodb.Alert.URI)
	if err != nil {
		logrus.Fatal(err)
	}

	// repository
	alertRepo := mongo.NewAlertRepository(
		mongoAlertClient, s.cfg.Adapter.Mongodb.Alert.Database, s.cfg.Adapter.Mongodb.Alert.Collection,
	)

	// use_case
	alertUseCase := usecase.NewAlertUseCase(s.cfg, alertRepo)
	
	// handler
	alertHandler := v1.NewAlertHandler(s.cfg, s.echo.Group(s.cfg.Http.BasePath), alertUseCase)
	alertHandler.MapRoutes()

	go func() {
		if err = s.runHttpServer(); err != nil {
			s.logger.Fatalln("Error starting Server:", err)
		}
	}()

	// Set up a channel to listen to interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Block on this channel listening for those previously defined syscalls assign
	// to variable, so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	<-ctx.Done()
	s.logger.Println("Server is shutting down due to", interrupt)
	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.logger.Fatalln("Server was unable to gracefully shutdown due to err:", err)
	}

	s.logger.Fatalln("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
