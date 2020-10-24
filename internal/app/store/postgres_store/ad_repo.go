package postgres_store

import (
	"database/sql"
	"fmt"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/MeguMan/buyer-exp-test/internal/app/store"
	"log"
)

type AdRepository struct {
	store *Store
}

func (r *AdRepository) Create(a *model.Ad) (int, error) {
	//VALIDATE HERE
	if existing, err := r.FindByLink(a.Link); existing != nil {
		return existing.ID, err
	}

	var id int
	a.Price = a.ParsePrice(a.Link)
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

func (r *AdRepository) CheckPrice() {
	rows, err := r.store.db.Query("SELECT * FROM ads")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var aa []model.Ad

	for rows.Next() {
		a := model.Ad{}
		err := rows.Scan(&a.ID, &a.Link, &a.Price)
		if err != nil {
			log.Print(err)
			continue
		}
		aa = append(aa, a)
	}

	for _, a := range aa {
		newPrice := a.ParsePrice(a.Link)

		if a.Price != newPrice {
			a.Price = newPrice
			r.UpdatePrices(&a)
		}
	}
}

func (r *AdRepository) UpdatePrices(a *model.Ad) {
	_, err := r.store.db.Exec("UPDATE ads SET price = $1 WHERE link = $2", a.Price, a.Link)
	if err != nil {
		panic(err)
	}
	fmt.Printf("У объявления по ссылку %s обновилась цена, теперь она равна %d \n", a.Link, a.Price)

	rows, err := r.store.db.Query(
		"SELECT email FROM users INNER JOIN users_ads ON (users.user_id = users_ads.user_id) WHERE users_ads.ad_id = $1;",
		a.ID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		u := model.User{}
		err := rows.Scan(&u.Email)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(u)
		err = u.SendEmail(a)
		if err != nil {
			log.Println(err)
		}
	}
}
