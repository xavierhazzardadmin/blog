package routes

import (
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
    if err := c.Validate(p); err != nil {
        return err
    }
    if err := db.Post(p); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, p)
}

//  Handler for getting individual posts

func Get(c echo.Context) error {
    id := c.Param("id")
    p, err := db.Get(id)

    if err != nil {
       return echo.NewHTTPError(http.StatusBadRequest, err.Error()) 
    }
    
    return c.JSON(http.StatusOK, p)
}

//  Handler for getting all author posts
func GetAll(c echo.Context) error {
    author := c.FormValue("author")
    posts, err :=  db.GetAll(author)

    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    return c.JSON(http.StatusOK, posts)

}

//  Handler for updating posts

//  Handler for getting titles of each post to cache

//  Handler for deleting posts
