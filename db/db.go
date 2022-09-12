package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/xavierhazzardadmin/blog/helpers"
	"github.com/xavierhazzardadmin/blog/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne(title string) (*models.Post, error) {
	var r bson.M
	var result models.Post

	user := helpers.GetEnv("MongoUser")
	pass := helpers.GetEnv("MongoPass")

	authURI := fmt.Sprintf("mongodb+srv://%s:%s@blog.w53zlvb.mongodb.net/?retryWrites=true&w=majority", user, pass)

	ctx := context.TODO()
	opts := options.Client().ApplyURI(authURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("blog")

	articles := db.Collection("articles")

	c := articles.FindOne(ctx, bson.M{"slug": strings.ToLower(title)})
	c.Decode(&result)

	rBytes, _ := bson.Marshal(r)
	bson.Unmarshal(rBytes, result)

	return &result, err
}
