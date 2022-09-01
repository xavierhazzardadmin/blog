package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/xavierhazzardadmin/blog/models"
)

const (
	projectId      string = "xaviers-blog"
	collectionName string = "posts"
)

func Save(post *models.Post) (*models.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Error creating a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"author":      post.Author,
		"content":     post.Content,
		"id":          post.ID,
		"metatitle":   post.MetaTitle,
		"publishedAt": post.Published,
		"updatedAt":   post.Updated,
		"slug":        post.Slug,
		"summary":     post.Summary,
	})

	if err != nil {
		log.Fatalf("Error adding a new post: %v", err)
		return nil, err
	}

	return post, err
}

func NewPost(post *models.Post) *models.Post {
	return &models.Post{}
}
