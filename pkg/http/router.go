package http

import (
	"net/http"

	"logs-hub-backend/pkg/http/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	adminHandler := handlers.NewAdminHandler()
	serviceHandler := handlers.NewServiceHandler()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	adminGroup := e.Group("/admin")
	serviceGroup := e.Group("/services")

	// Admin routes
	adminGroup.POST("/register", adminHandler.Register)
	adminGroup.POST("/login", adminHandler.Login)
	adminGroup.POST("/logout", adminHandler.Logout)

	// Service routes
	serviceGroup.POST("/register", serviceHandler.Register)
	serviceGroup.GET("/apikey", serviceHandler.GetAPIKey)
	serviceGroup.GET("", serviceHandler.GetServices)
	// serviceGroup.PUT("/:id", serviceHandler.UpdateService)
	serviceGroup.DELETE("/:id", serviceHandler.DeleteService)
	serviceGroup.GET("/:id/logs", serviceHandler.GetLog)
	serviceGroup.POST("/:id/logs/filter", serviceHandler.GetLogByFilter)

	return e
}
