package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllArticles(c echo.Context) error {
	articles := "This is all articles"
	return c.String(http.StatusOK, articles)
}

func GetNewArticles(c echo.Context) error {
	return c.String(http.StatusOK, "Newest Article")
}

func GetArticle(c echo.Context) error {
	return c.String(http.StatusOK, "Single Article")
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
