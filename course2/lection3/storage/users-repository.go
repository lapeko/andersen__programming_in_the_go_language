package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/internal/app/models"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
)

const usersTableName string = "users"

type UsersRepository struct {
	db *sql.DB
}

func (u *UsersRepository) CreateUser(user *models.User) (*models.User, error) {
	log := logger.Get()

	r, err := u.GetByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if r != nil {
		return nil, nil
	}

	query := fmt.Sprintf("INSERT INTO %s (email, password) VALUES ($1, $2) RETURNING id", usersTableName)

	err = u.db.QueryRow(query, user.Email, user.Password).Scan(&user.Id)

	if err != nil {
		log.Error("Create user sql query error")
		return nil, err
	}

	return user, nil
}

func (u *UsersRepository) GetAll() ([]*models.User, error) {
	log := logger.Get()
	query := fmt.Sprintf("SELECT * FROM %s", usersTableName)
	r, err := u.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := r.Close()
		if err != nil {
			log.Error(err)
		}
	}()

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

func (u *UsersRepository) GetByEmail(email string) (*models.User, error) {
	log := logger.Get()
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", usersTableName)
	user := &models.User{}
	err := u.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}

	return user, nil
}
