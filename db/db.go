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
        log.Info("Mongo Connection Error:")
        return err
	}
	defer client.Disconnect(ctx)

	db := client.Database("blog")
	articles := db.Collection("articles")

    _, err = articles.InsertOne(ctx, post)
    if err != nil {
        log.Info("Mongo Insertion Error:")
        return err
    }

    return err
}

func Get(id string) (*models.Post, error) {
	var r bson.M
	var result models.Post
    objID, err := primitive.ObjectIDFromHex(id)

    if err != nil {
        log.Info("Mongo ID Error:")
        return nil, err
    }

	ctx := context.TODO()
	opts := options.Client().ApplyURI(authURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
        log.Info("Mongo Connection Error")
        return nil, err
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
        log.Info("Mongo Connection Error:")
        return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database("blog")
	articles := db.Collection("articles")

    articleCursor, err := articles.Find(ctx, bson.M{"author": author})
    if err != nil {
        log.Info("Mongo Cursor Error:")
        return nil, err
    }

    if err = articleCursor.All(ctx, &posts); err != nil {
        log.Info("Mongo Binding Error:")
        return nil, err
    }

    return posts, err
}


func Delete(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)

    if err != nil {
        log.Info("Mongo ID Error:")
        return err
    }

    ctx := context.TODO()
    opts := options.Client().ApplyURI(authURI)

    client, err  := mongo.Connect(ctx, opts)
    if err != nil {
        log.Info("Mongo Connection Error:")
        return err
    }
    defer client.Disconnect(ctx)

    db := client.Database("blog")
    articles := db.Collection("articles")

    _, err = articles.DeleteOne(ctx, bson.M{"_id": objID} )

    return err
}

func Update(id string, post *models.Post) error {
    objID, err := primitive.ObjectIDFromHex(id)

    if err != nil {
        log.Info("Mongo ID Error:")
        return err
    }
    ctx := context.TODO()
    opts := options.Client().ApplyURI(authURI)

    client, err := mongo.Connect(ctx, opts)

    if err != nil {
        log.Info("Mongo Connection Error:")
        return err 
    }
    defer client.Disconnect(ctx)

    db := client.Database("blog")
    articles := db.Collection("articles")

    _, err = articles.ReplaceOne(ctx, bson.D{primitive.E{Key: "_id", Value: objID}}, post)

    if err != nil {
        log.Info("Mongo Replace Error:")
        return err
    }

    return err
}

func GetCache() ([]*models.Post, error) {
    posts := make([]*models.Post, 10)

    ctx := context.TODO()
    opts := options.Client().ApplyURI(authURI)

    client, err := mongo.Connect(ctx, opts)

    if err != nil {
        log.Info("Mongo Connection Error:")
        return nil, err
    }
    defer client.Disconnect(ctx)

    db := client.Database("blog")
    articles := db.Collection("articles")

    filter := bson.D{}
    findOpts := options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: -1}}).SetLimit(10)

    cursor, err := articles.Find(ctx, filter, findOpts)

    if err != nil {
        log.Info("Mongo Find Error:")
        return nil, err
    }

    if err = cursor.All(ctx, &posts); err != nil {
        log.Info("Mongo Cursor Error:")
        return nil, err
    }

    return posts, err
}
