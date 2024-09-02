package storage

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
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
