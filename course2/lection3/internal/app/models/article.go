package models

type Article struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	AuthorId uint   `json:"authorId"`
	Content  string `json:"content"`
}
