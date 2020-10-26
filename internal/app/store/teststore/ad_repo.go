package teststore

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
)

type AdRepository struct {
	store *Store
	ads   map[int]*model.Ad
}

func (r *AdRepository) FindByLink(s string) (*model.Ad, error) {
	panic("implement me")
}

func (r *AdRepository) CheckPrice() {
	panic("implement me")
}

func (r *AdRepository) UpdatePrices(a *model.Ad) error {
	panic("implement me")
}

func (r *AdRepository) Create(a *model.Ad) error {
	if err := a.Validate(); err != nil {
		return err
	}

	return nil
}
