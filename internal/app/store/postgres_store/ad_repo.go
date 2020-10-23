package postgres_store

import (
	"database/sql"
	"fmt"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/MeguMan/buyer-exp-test/internal/app/store"
)

type AdRepository struct {
	store *Store
}

func (r *AdRepository) Create(a *model.Ad) (int,error) {
	//VALIDATE HERE
	if existing, err := r.FindByLink(a.Link); existing != nil {
		return existing.ID, err
	}

	var id int
	err := r.store.db.QueryRow("INSERT INTO ads (link, price) VALUES ($1, $2) returning ad_id",
		a.Link, a.Price).Scan(&id)
	fmt.Println(id)
	return id, err
}

func (r *AdRepository) FindByLink(link string) (*model.Ad, error) {
	a := &model.Ad{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM ads WHERE link = $1",
		link,
	).Scan(
		&a.ID,
		&a.Link,
		&a.Price,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return a, nil
}