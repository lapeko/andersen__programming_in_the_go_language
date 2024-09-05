package storage

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
)

type Storage struct {
	config             *Config
	db                 *sql.DB
	usersRepository    *UsersRepository
	articlesRepository *ArticlesRepository
}

func New(c *Config) *Storage {
	return &Storage{
		config: c,
	}
}

func (s *Storage) Open() (err error) {
	if s.db != nil {
		return
	}

	db, err := sql.Open("pgx", s.config.DBUri)
	if err != nil {
		return
	}

	err = db.Ping()
	if err == nil {
		s.db = db
	}
	return
}

func (s *Storage) Close() (err error) {
	err = s.db.Close()
	return
}

func (s *Storage) Users() *UsersRepository {
	log := logger.Get()
	if s.db == nil {
		log.Fatalln("DB is not initialized")
	}
	if s.usersRepository == nil {
		s.usersRepository = &UsersRepository{s.db}
	}
	return s.usersRepository
}

func (s *Storage) Articles() *ArticlesRepository {
	log := logger.Get()
	if s.db == nil {
		log.Fatalln("DB is not initialized")
	}
	if s.articlesRepository == nil {
		s.articlesRepository = &ArticlesRepository{s.db}
	}
	return s.articlesRepository
}
