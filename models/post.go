package models

type Post struct {
	Author    string `json:"author"`
	Content   string `json:"content"`
	ID        int    `json:"_id" `
	MetaTitle string `json:"metaTitle" `
	Published string `json:"published"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Updated   string `json:"updated"`
	Summary   string `json:"summary"`
}
