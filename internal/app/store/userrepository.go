package store

import "github.com/mmaxim2710/firstWebApp/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	r.store.logger.Debugf("Creating user with email=%s", u.Email)
	err := u.Validate()
	if err != nil {
		return nil, err
	}

	err = u.BeforeCreate()
	if err != nil {
		return nil, err
	}

	err = r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	r.store.logger.Debugf("Querying user with email=%s", email)
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1", email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return u, nil
}
