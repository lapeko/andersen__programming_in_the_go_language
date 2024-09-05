package storage

import (
	"database/sql"
	"fmt"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/internal/app/models"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
)

const usersTableName string = "users"

type UsersRepository struct {
	db *sql.DB
}

func (u *UsersRepository) GetAll() ([]*models.User, error) {
	log := logger.Get()
	query := fmt.Sprintf("SELECT * FROM %s", usersTableName)
	r, err := u.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer func(r *sql.Rows) {
		err := r.Close()
		if err != nil {
			log.Error(err)
		}
	}(r)

	users := make([]*models.User, 0)

	for r.Next() {
		user := &models.User{}
		err := r.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			log.Warningln("Parse user failure", err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
