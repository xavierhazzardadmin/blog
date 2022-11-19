package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xavierhazzardadmin/blog/db"
	"github.com/xavierhazzardadmin/blog/models"
)

func Post(c echo.Context) error {
    p := new(models.Post)
    if err := c.Bind(p); err != nil {
        fmt.Println("Routes 1")
        return err
    }
    if err := c.Validate(p); err != nil {
        fmt.Println("Routes 2")
        return err
    }
    if err := db.Post(p); err != nil {
        fmt.Println("Routes 3")
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
    p := new(models.Post)

    if err := c.Bind(p); err != nil {
        fmt.Println("Routes 1")
        return err
    }

    if err := c.Validate(p); err != nil {
        fmt.Println("Routes 2")
        return err
    }

    if err := db.Update(id, p); err != nil {
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

func Cache(c echo.Context) error {
    posts, err := db.GetCache()

    if err != nil{
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    fmt.Println(len(posts))

    return c.JSON(http.StatusOK, posts)
}
