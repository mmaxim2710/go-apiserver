package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
	"github.com/sirupsen/logrus"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	logger         *logrus.Logger
}

func New(db *sql.DB, logger *logrus.Logger) *Store {
	logger.Debug("Initializing new Store")
	return &Store{
		db:     db,
		logger: logger,
	}
}

func (s *Store) User() store.UserRepository {
	s.logger.Debug("Initializing User Repository")
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}
