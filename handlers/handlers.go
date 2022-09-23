package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavierhazzardadmin/blog/db"
	"github.com/xavierhazzardadmin/blog/models"
)

func SaveHandler(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err})
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	Post, err := db.Save(&post)

	if err != nil {
		log.Fatalf("error saving post: %v", err)
		fmt.Println(err, Post)
	}
}

func FindHandler(c *gin.Context) {
	var search models.Search
	if err := c.BindJSON(search); err != nil {
		c.JSON(400, gin.H{"error": err})
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err := db.Find()
}
