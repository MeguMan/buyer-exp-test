package store

import "github.com/MeguMan/buyer-exp-test/internal/app/model"

type UserRepository interface {
	Create(user *model.User) (int, error)
	FindByEmail(string) (*model.User, error)
}