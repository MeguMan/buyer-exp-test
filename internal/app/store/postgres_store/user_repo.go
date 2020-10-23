package postgres_store

import (
	"database/sql"
	"fmt"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/MeguMan/buyer-exp-test/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func(r *UserRepository) Create(u *model.User) (int, error) {
	//Need to validate
	if existing, err := r.FindByEmail(u.Email); existing != nil {
		return existing.ID, err
	}

	var id int
	err := r.store.db.QueryRow("INSERT INTO users (email) VALUES ($1) returning user_id",
		u.Email).Scan(&id)
	fmt.Println(id)
	return id, err
}

func(r *UserRepository) FindByEmail(email string) (*model.User, error){
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
