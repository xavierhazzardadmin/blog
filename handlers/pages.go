package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/xavierhazzardadmin/blog/db"
)

func GetAllArticles(c echo.Context) error {
	articles := "This is all articles"
	return c.String(http.StatusOK, articles)
}

func GetNewArticles(c echo.Context) error {
	return c.String(http.StatusOK, "Newest Article")
}

func GetArticle(c echo.Context) error {
	res, err := db.FindOne(strings.ToLower(c.Param("title")))

	if err != nil {
		c.Logger().Panic(err)
	}
	return c.String(http.StatusOK, fmt.Sprintf("You are viewing the article: %s", res.Title))
}

func GetProfile(c echo.Context) error {
	name := "Xavier"

	if c.Param("name") != "" {
		name = c.Param("name")
	}
	return c.String(http.StatusOK, "Hello "+name+" Welcome to your profile.")

}

func GetAllProfiles(c echo.Context) error {
	return c.String(http.StatusOK, "All Profiles")
}
