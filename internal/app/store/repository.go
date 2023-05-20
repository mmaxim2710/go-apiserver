package store

import "github.com/mmaxim2710/firstWebApp/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
