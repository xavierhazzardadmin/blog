package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
    Author    string `json:"author" bson:"author"`
    Content   string `json:"content" bson:"content"`
    ID        primitive.ObjectID    `json:"_id" bson:"_id"`
    MetaTitle string `json:"metaTitle" bson:"metaTitle"`
    Published string `json:"published" bson:"published"`
    Slug      string `json:"slug" bson:"slug"`
    Title     string `json:"title" bson:"title"`
    Updated   string `json:"updated" bson:"updated"`
    Summary   string `json:"summary" bson:"summary"`
}
