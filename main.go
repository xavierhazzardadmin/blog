package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xavierhazzardadmin/blog/routes"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/posts", routes.Get)
    e.POST("/posts", routes.Post)

	e.Logger.Fatal(e.Start(":8080"))

}
