package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
)

type Store struct {
	config         *config.Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(c *config.Config) *Store {
	return &Store{
		config: c,
	}
}
func (s *Store) Open() error {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.config.DB.Host, s.config.DB.Port, s.config.DB.User, s.config.DB.Password, s.config.DB.DBName))
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}
