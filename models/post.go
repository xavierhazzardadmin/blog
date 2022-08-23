package main

type Post struct {
	Author    string `json:"author"`
	Content   string `json:"content"`
	ID        int    `json:"id"`
	MetaTitle string `json:"metaTitle"`
	Published int64  `json:"published"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Updated   int64  `json:"updated"`
}
