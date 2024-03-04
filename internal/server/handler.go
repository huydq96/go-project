package server

import (
	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	//// repos
	//userRepo := authRepository.NewUserRepository(s.db)
	//todoRepo := todoRepository.NewTodoRepository(s.db)
	//
	////usecase
	//authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	//todoUC := todoUsecase.NewTodoUseCase(todoRepo, userRepo)
	//
	////handler
	//authHandler := authHttp.NewAuthHandler(authUC)
	//todoHandler := todoHttp.NewTodoHandler(todoUC)
	//
	////middlewares
	//mw := middlewares.NewMiddlewareManager(authUC)
	//
	//e.Use(middleware.BodyLimit("2M"))
	//// e.Use(mw.JWTValidation())
	//
	////versioning
	//v1 := e.Group("/api/v1")
	//
	//authGroup := v1.Group("/auth")
	//todoGroup := v1.Group("/todos")
	//
	//authHttp.MapAuthRoutes(authGroup, authHandler)
	//todoHttp.MapAuthRoutes(todoGroup, todoHandler, mw)

	return nil
}
