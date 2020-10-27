package teststore

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (model.User, error) {
	return *u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	panic("implement me")
}
