package store

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/firstWebApp/internal/app/model"
)

type UserRepository interface {
	Create(user *model.User) error
	Find(uuid uuid.UUID) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
