package postgres_store

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (model.User, error) {
	if u, err := r.FindByEmail(u.Email); u != nil {
		return *u, err
	}

	err := r.store.db.QueryRow("INSERT INTO users (email) VALUES ($1) returning user_id",
		u.Email).Scan(&u.ID)

	return *u, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email); err != nil {
		return nil, err
	}

	return u, nil
}
