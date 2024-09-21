package repository

import (
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/internal/models"
	"gorm.io/gorm"
)

type Repo struct {
	Articles *ArticlesRepo
}

type ArticlesRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{Articles: &ArticlesRepo{db: db}}
}

func (r *ArticlesRepo) Create(a *models.Article) error {
	return r.db.Create(a).Error
}

func (r *ArticlesRepo) GetAll(a *[]models.Article) error {
	return r.db.Find(a).Error
}

func (r *ArticlesRepo) GetById(a *models.Article, id string) error {
	return r.db.Find(a, id).Error
}

func (r *ArticlesRepo) DeleteById(id string) error {
	return r.db.Delete(&models.Article{}, id).Error
}

func (r *ArticlesRepo) Update(a *models.Article) error {
	return r.db.Save(a).Error
}
