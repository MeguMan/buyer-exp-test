package store

import "github.com/MeguMan/buyer-exp-test/internal/app/model"

type UserRepository interface {
	Create(u *model.User) error
	FindByEmail(string) (*model.User, error)
}
