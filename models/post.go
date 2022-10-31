package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
<<<<<<< HEAD
	Author    string `json:"author"`
	Content   string `json:"content"`
	ID        int    `json:"id"`
	MetaTitle string `json:"metaTitle"`
	Published string `json:"published"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Updated   string `json:"updated"`
	Summary   string `json:"summary"`
=======
    Author    string `json:"author" bson:"author"`
    Content   string `json:"content" bson:"content"`
    ID        primitive.ObjectID    `json:"_id" bson:"_id"`
    MetaTitle string `json:"metaTitle" bson:"metaTitle"`
    Published string `json:"published" bson:"published"`
    Slug      string `json:"slug" bson:"slug"`
    Title     string `json:"title" bson:"title"`
    Updated   string `json:"updated" bson:"updated"`
    Summary   string `json:"summary" bson:"summary"`
>>>>>>> feature/gin-to-echo
}

type Search struct {
	QueryString string `json:"queryString"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
}
