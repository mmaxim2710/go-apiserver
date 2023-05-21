package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mmaxim2710/firstWebApp/internal/app/logger"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	logger.GetLogger().Debug("Initializing new Store")
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	logger.GetLogger().Debug("Initializing User Repository")
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}
