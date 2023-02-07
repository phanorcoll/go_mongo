package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App struct    Holds sharable information about the server
type App struct {
	server *echo.Echo
}

// New function    creates a new Echo server with configured middlaware
func New() *App {
	server := echo.New()

	//middleware
	server.Use(middleware.Recover())

	return &App{
		server: server,
	}
}

// ConfigureRoutes method    creates endpoints.
func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
}

// Start method    initiates the Echo server
func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(":5000")
}
