package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/xavierhazzardadmin/blog/helpers"
	"github.com/xavierhazzardadmin/blog/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var authURI string = fmt.Sprintf("mongodb+srv://%s:%s@blog.w53zlvb.mongodb.net/?retryWrites=true&w=majority", helpers.GetEnv("MongoUser"), helpers.GetEnv("MongoPass"))


func Insert(post *models.Post) *mongo.InsertOneResult {
    post.ID = primitive.NewObjectID()
	ctx := context.TODO()
	opts := options.Client().ApplyURI(authURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("blog")
	articles := db.Collection("articles")

    result, err := articles.InsertOne(ctx, &post)
    if err != nil {
        panic(err)
    }

    return result
}

func Get(id string) (*models.Post, error) {
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
		panic(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("blog")
	articles := db.Collection("articles")

    articleCursor, err := articles.Find(ctx, bson.M{"author": strings.ToLower(author)})
    if err != nil {
        panic(err)
    }

    if err = articleCursor.All(ctx, &posts); err != nil {
       panic(err)
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
