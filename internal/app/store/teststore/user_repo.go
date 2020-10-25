package teststore

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) (int, error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}

	return u.ID, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	panic("implement me")
}
