package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xavierhazzardadmin/blog/db"
	"github.com/xavierhazzardadmin/blog/models"
)

//  Handler for inserting posts

func Post(c echo.Context) error {
    p := new(models.Post)
    if err := c.Bind(p); err != nil {
        return &echo.BindingError{}
    }
    fmt.Println(p)
    if err := db.Post(p); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, p)
}

//  Handler for getting individual posts

func Get(c echo.Context) error {
    q := new(models.Query)
    if err := c.Bind(q); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    p, err := db.Get(q.ID)

    if err != nil {
       return echo.NewHTTPError(http.StatusBadRequest, err.Error()) 
    }
    
    return c.JSON(http.StatusOK, p)
}

//  Handler for getting all author posts

//  Handler for updating posts

//  Handler for getting titles of each post to cache

//  Handler for deleting posts
