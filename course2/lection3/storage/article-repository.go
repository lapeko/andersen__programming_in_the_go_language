package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/internal/app/models"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
	"github.com/sirupsen/logrus"
)

const articlesTableName = "articles"

type ArticlesRepository struct {
	db *sql.DB
}

func (a *ArticlesRepository) GetAll() ([]*models.Article, error) {
	log := logger.Get()

	query := fmt.Sprintf("SELECT * FROM %s", articlesTableName)
	r, err := a.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer func(r *sql.Rows) {
		err := r.Close()
		if err != nil {
			logrus.Errorln(err)
		}
	}(r)

	articles := make([]*models.Article, 0)

	for r.Next() {
		article := &models.Article{}
		err = r.Scan(&article.Id, &article.Content, &article.AuthorId, &article.Title)

		if err != nil {
			log.Warningln(err)
			continue
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (a *ArticlesRepository) GetArticleById(id int) (*models.Article, error) {
	query := fmt.Sprintf("SELECT * from %s WHERE id = $1", articlesTableName)

	article := &models.Article{}

	err := a.db.QueryRow(query, id).Scan(&article.Id, &article.Title, &article.AuthorId, &article.Content)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return article, nil
}
