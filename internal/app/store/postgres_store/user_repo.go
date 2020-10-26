package postgres_store

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if exists, err := r.FindByEmail(u.Email); exists != nil {
		return err
	}

	err := r.store.db.QueryRow("INSERT INTO users (email) VALUES ($1) returning user_id",
		u.Email).Scan(&u.ID)

	return err
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
