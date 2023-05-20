package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
	"github.com/sirupsen/logrus"
)

type Store struct {
	config         *config.Config
	db             *sql.DB
	userRepository *UserRepository
	logger         *logrus.Logger
}

func New(c *config.Config, logger *logrus.Logger) *Store {
	logger.Debugf("Initializating new Store with params: host=%s, port=%s, user=%s, db_name=%s",
		c.DB.Host, c.DB.Port, c.DB.User, c.DB.DBName)
	return &Store{
		config: c,
		logger: logger,
	}
}
func (s *Store) Open() error {
	s.logger.Debugf("Opening connection with db \"%s\"", s.config.DB.DBName)
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
	s.logger.Debugf("Closing connection with db \"%s\"", s.config.DB.DBName)
	err := s.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) User() *UserRepository {
	s.logger.Debug("Initializing User Repository")
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}
