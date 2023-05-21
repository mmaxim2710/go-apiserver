package teststore

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/firstWebApp/internal/app/model"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[uuid.UUID]*model.User
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

	newUuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	u.UUID = newUuid
	r.users[u.UUID] = u
	return nil
}

func (r *UserRepository) Find(uuid uuid.UUID) (*model.User, error) {
	u, ok := r.users[uuid]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
