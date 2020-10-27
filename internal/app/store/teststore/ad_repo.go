package teststore

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/emailsender"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
)

type AdRepository struct {
	store *Store
}

func (r *AdRepository) FindByLink(s string) (*model.Ad, error) {
	panic("implement me")
}

func (r *AdRepository) CheckPrice() {
	panic("implement me")
}

func (r *AdRepository) UpdatePrices(a *model.Ad, sender *emailsender.Sender) error {
	panic("implement me")
}

func (r *AdRepository) Create(a *model.Ad) (model.Ad, error) {
	return *a, nil
}
