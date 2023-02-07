package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phanorcoll/go_mongo/pkg/config"
	"github.com/phanorcoll/go_mongo/pkg/data"
	"github.com/phanorcoll/go_mongo/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// App struct    Holds sharable information about the server
type App struct {
	server  *echo.Echo
	userSvc services.IUserService
	cfg     *config.Settings
}

// New function    creates a new Echo server with configured middlaware
func New(cfg *config.Settings, client *mongo.Client) *App {
	server := echo.New()

	//middleware
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	//providers
	userProvider := data.NewUserProvider(cfg, client)

	//seervices
	userSvc := services.NewUserService(cfg, userProvider)

	return &App{
		server:  server,
		userSvc: userSvc,
		cfg:     cfg,
	}
}

// ConfigureRoutes method    creates endpoints.
func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
	a.server.POST("/v1/public/account/register", a.Register)
	a.server.POST("/v1/public/account/login", a.Login)

	protected := a.server.Group("vi/api")

	middlaware := Middleware{config: a.cfg}

	protected.Use(middlaware.Auth)
	protected.GET("/secret", func(c echo.Context) error {
		userId := c.Get("user").(string)
		return c.String(200, userId)
	})
}

// Start method    initiates the Echo server
func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(":5000")
}
