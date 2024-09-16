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

func (a *ArticlesRepository) GetById(id int) (*models.Article, error) {
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

func (a *ArticlesRepository) Create(article *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author_id, content) VALUES ($1, $2, $3) RETURNING id", articlesTableName)
	err := a.db.QueryRow(query, article.Title, article.AuthorId, article.Content).Scan(&article.Id)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a *ArticlesRepository) Delete(id int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", articlesTableName)
	res, err := a.db.Exec(query, id)

	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (a *ArticlesRepository) Update(article *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("UPDATE %s SET title=$1, author_id=$2, content=$3 WHERE id=$4", articlesTableName)
	r, err := a.db.Exec(query, article.Title, article.AuthorId, article.Content, article.Id)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := r.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, nil
	}

	return article, nil
}
