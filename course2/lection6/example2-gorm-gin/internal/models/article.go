package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `form:"title"`
	Content string `form:"content"`
	Author  string `form:"author"`
}
