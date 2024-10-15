package server

import (
	"github.com/labstack/echo/v4"
	"github.com/realiza-org/go-service/internal/task"
)

func NewServer() *echo.Echo {
	e := echo.New()

	// Initialize task handlers and routes
	taskGroup := e.Group("/tasks")
	task.RegisterRoutes(taskGroup)

	return e
}
