package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type( 
    Post struct {
        Author    string `json:"author" bson:"author" validate:"required"`
        Content   string `json:"content" bson:"content" validate:"required"`
        ID        primitive.ObjectID    `json:"_id" bson:"_id"`
        MetaTitle string `json:"metaTitle" bson:"metaTitle" validate:"required"`
        Published string `json:"published" bson:"published" validate:"required"`
        Slug      string `json:"slug" bson:"slug" validate:"required"`
        Title     string `json:"title" bson:"title" validate:"required"`
        Updated   string `json:"updated" bson:"updated" validate:"required"`
        Summary   string `json:"summary" bson:"summary" validate:"required"`
    }

    Author struct {
        Name     string 
        Age      string
        Email    string
        Color    string
        Location string
    }


)

