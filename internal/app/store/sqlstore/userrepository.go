package sqlstore

import (
	"database/sql"
	"github.com/mmaxim2710/firstWebApp/internal/app/logger"
	"github.com/mmaxim2710/firstWebApp/internal/app/model"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	logger.GetLogger().Debugf("Creating user with email=%s", u.Email)
	err := u.Validate()
	if err != nil {
		return err
	}

	err = u.BeforeCreate()
	if err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING uuid",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.UUID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	logger.GetLogger().Debugf("Querying user with email=%s", email)
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT uuid, email, encrypted_password FROM users WHERE email = $1", email,
	).Scan(&u.UUID, &u.Email, &u.EncryptedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
