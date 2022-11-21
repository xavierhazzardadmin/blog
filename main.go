package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xavierhazzardadmin/blog/models"
	"github.com/xavierhazzardadmin/blog/routes"
)

func main() {
	e := echo.New()
    e.Validator = models.NewValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
    e.GET("/posts/:id", routes.Get)
    e.POST("/posts", routes.Post)
    e.DELETE("/posts/:id", routes.Delete)
    e.PUT("/posts/:id", routes.Update)
    e.POST("/posts/q", routes.GetAll)
    e.GET("/cache", routes.Cache)

	e.Logger.Fatal(e.Start(":8080"))
}
