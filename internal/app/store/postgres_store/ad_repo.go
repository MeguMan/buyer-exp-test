package postgres_store

import (
	"encoding/json"
	"fmt"
	"github.com/MeguMan/buyer-exp-test/internal/app/emailsender"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"log"
	"os"
)

type AdRepository struct {
	store *Store
}

func (r *AdRepository) Create(a *model.Ad) (model.Ad, error) {
	if a, err := r.FindByLink(a.Link); a != nil {
		return *a, err
	}
	a.Price = a.ParsePrice(a.Link)
	err := r.store.db.QueryRow("INSERT INTO ads (link, price) VALUES ($1, $2) returning ad_id",
		a.Link, a.Price).Scan(&a.ID)

	return *a, err
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
		return nil, err
	}

	return a, nil
}

func (r *AdRepository) CheckPrice() {
	emailConfig := emailsender.NewConfig()
	configFile, err := os.Open("configs/db.json")
	if err != nil {
		log.Print(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&emailConfig); err != nil {
		log.Print(err.Error())
	}

	rows, err := r.store.db.Query("SELECT * FROM ads")
	if err != nil {
		log.Print(err)
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
			if err = r.UpdatePrices(&a, emailConfig); err != nil {
				log.Print(err)
			}
		}
	}
}

func (r *AdRepository) UpdatePrices(a *model.Ad, emailConfig *emailsender.Config) error {
	_, err := r.store.db.Exec("UPDATE ads SET price = $1 WHERE link = $2", a.Price, a.Link)
	if err != nil {
		return err
	}
	fmt.Printf("У объявления по ссылку %s обновилась цена, теперь она равна %d \n", a.Link, a.Price)

	rows, err := r.store.db.Query(
		"SELECT email FROM users INNER JOIN users_ads ON (users.user_id = users_ads.user_id) WHERE users_ads.ad_id = $1;",
		a.ID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		u := &model.User{}
		err := rows.Scan(&u.Email)
		if err != nil {
			log.Println(err)
			continue
		}

		err = emailsender.New(emailConfig).SendEmail(u, a)
		if err != nil {
			return err
		}
	}

	return nil
}
