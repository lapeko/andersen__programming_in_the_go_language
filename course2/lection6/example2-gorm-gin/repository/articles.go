package repository

import "gorm.io/gorm"

type ArticlesRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ArticlesRepo {
	return &ArticlesRepo{db: db}
}
