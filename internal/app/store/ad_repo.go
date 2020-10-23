package store

import "github.com/MeguMan/buyer-exp-test/internal/app/model"

type AdRepository interface {
	Create(ad *model.Ad) (int, error)
	FindByLink(string) (*model.Ad, error)
}

