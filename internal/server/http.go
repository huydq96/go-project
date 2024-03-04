package server

import (
	"fmt"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "10M"
)

func (s *Server) runHttpServer() error {
	s.mapRoutes()

	s.echo.Server.ReadTimeout = s.cfg.Server.Timeout.Read
	s.echo.Server.WriteTimeout = s.cfg.Server.Timeout.Write
	s.echo.Server.IdleTimeout = s.cfg.Server.Timeout.Idle
	s.echo.Server.MaxHeaderBytes = maxHeaderBytes

	return s.echo.Start(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port))
}

func (s *Server) mapRoutes() {
	s.echo.Logger.SetLevel(log.DEBUG)
	s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.BodyLimit(bodyLimit))
}
