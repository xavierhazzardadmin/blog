package db

import (
	"context"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/xavierhazzardadmin/blog/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var authURI string = fmt.Sprintf("%s", os.Getenv("MongoURI"))


func Post(post *models.Post) error {
    post.ID = primitive.NewObjectID()
	ctx := context.TODO()
	opts := options.Client().ApplyURI(authURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
        fmt.Println("here", err.Error())
	}
	defer client.Disconnect(ctx)

	db := client.Database("blog")
	articles := db.Collection("articles")

    _, err = articles.InsertOne(ctx, &post)
    if err != nil {
        panic(err)
    }

    return err
}

func Get(id string) (*models.Post, error) {
    fmt.Println(authURI)
	var r bson.M
	var result models.Post
    objID, err := primitive.ObjectIDFromHex(id)

    if err != nil {
        panic(err)
    }

	ctx := context.TODO()
	opts := options.Client().ApplyURI(authURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

    db := client.Database("blog")
	articles := db.Collection("articles")

    c := articles.FindOne(ctx, bson.M{"_id": objID})
	c.Decode(&result)

	rBytes, _ := bson.Marshal(r)
	bson.Unmarshal(rBytes, result)

	return &result, err
}


func GetAll(author string) ([]*models.Post, error) {
    var posts []*models.Post
    
	ctx := context.TODO()
	opts := options.Client().ApplyURI(authURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
        log.Info("Mongo connection error")
		panic(err.Error())
	}
	defer client.Disconnect(ctx)

	db := client.Database("blog")
	articles := db.Collection("articles")

    articleCursor, err := articles.Find(ctx, bson.M{"author": author})
    if err != nil {
        log.Info("Mongo cursor error")
        panic(err.Error())
    }

    if err = articleCursor.All(ctx, &posts); err != nil {
       log.Info("Mongo Binding issue")
       panic(err.Error())
    }

    return posts, err
}


func Delete(id string) *mongo.DeleteResult {
    objID, err := primitive.ObjectIDFromHex(id)

    if err != nil {
        panic(err)
    }

    ctx := context.TODO()
    opts := options.Client().ApplyURI(authURI)

    client, err  := mongo.Connect(ctx, opts)
    if err != nil {
        panic(err)
    }
    defer client.Disconnect(ctx)

    db := client.Database("blog")
    articles := db.Collection("articles")

    result, err := articles.DeleteOne(ctx, bson.M{"_id": objID} )
    if err != nil {
        panic(err)
    }

    return result
}

func Update(post *models.Post) {
    ctx := context.TODO()
    opts := options.Client().ApplyURI(authURI)

    client, err := mongo.Connect(ctx, opts)

    if err != nil {
        panic(err)
    }
    defer client.Disconnect(ctx)

    db := client.Database("blog")
    articles := db.Collection("articles")

    _, err = articles.UpdateByID(ctx, post.ID, bson.D{primitive.E{Key:"title", Value: post.Title},})

    if  err != nil {
        panic(err)
    }
}
