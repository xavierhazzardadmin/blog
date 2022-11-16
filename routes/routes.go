package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xavierhazzardadmin/blog/db"
	"github.com/xavierhazzardadmin/blog/models"
)

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

func Get(c echo.Context) error {
    id := c.Param("id")
    p, err := db.Get(id)

    if err != nil {
       return echo.NewHTTPError(http.StatusBadRequest, err.Error()) 
    }
    
    return c.JSON(http.StatusOK, p)
}

func GetAll(c echo.Context) error {
    author := c.FormValue("author")
    posts, err :=  db.GetAll(author)

    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    return c.JSON(http.StatusOK, posts)

}

func Update(c echo.Context) error {
    id := c.Param("id")
    title := c.FormValue("title") 
    if err := db.Update(id, title); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, models.Response{Message: "Title updated successfully."})

}

//  Handler for getting titles of each post to cache

func Delete(c echo.Context) error {
    id := c.Param("id")

    if err := db.Delete(id); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, models.Response{Message: "Post deleted successfully."})
}
