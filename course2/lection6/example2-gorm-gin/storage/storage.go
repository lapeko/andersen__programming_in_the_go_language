package storage

import (
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/internal/models"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Storage struct {
	dsn  string
	db   *gorm.DB
	Repo *repository.Repo
}

func New(dsn string) *Storage {
	return &Storage{dsn: dsn}
}

func (s *Storage) Init() (err error) {
	if s.db != nil {
		log.Printf("Warning. An attempt to reinitialize the Storage")
		return nil
	}

	s.db, err = gorm.Open(postgres.Open(s.dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	s.Repo = repository.New(s.db)

	return
}

func (s *Storage) Migrate() error {
	return s.db.AutoMigrate(&models.Article{})
}

func (s *Storage) Close() {
	s.Close()
}
