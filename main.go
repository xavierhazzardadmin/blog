package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xavierhazzardadmin/blog/handlers"
	"github.com/xavierhazzardadmin/blog/helpers"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.GetNewArticles)
	e.GET("/article/:title", handlers.GetArticle)
	e.GET("/articles", handlers.GetAllArticles)
	e.GET("/profile", handlers.GetProfile)
	e.GET("/profile/:name", handlers.GetProfile)
	e.GET("/profiles", handlers.GetAllProfiles)

	fmt.Println("Server running on Port 8080")
	fmt.Println(helpers.SplitRows("This is a string\n\nThis is another string"))

	e.Logger.Fatal(e.Start(":8080"))

}

func insert() {
	fmt.Println("This is rocket league")
}
