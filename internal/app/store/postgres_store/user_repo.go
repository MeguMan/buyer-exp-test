package postgres_store

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (int, error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}

	if exists, err := r.FindByEmail(u.Email); exists != nil {
		return exists.ID, err
	}

	var id int
	err := r.store.db.QueryRow("INSERT INTO users (email) VALUES ($1) returning user_id",
		u.Email).Scan(&id)
	return id, err
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
