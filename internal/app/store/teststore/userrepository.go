package teststore

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/firstWebApp/internal/app/model"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	err := u.Validate()
	if err != nil {
		return err
	}

	err = u.BeforeCreate()
	if err != nil {
		return err
	}

	r.users[u.Email] = u
	newUuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	u.UUID = newUuid
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
