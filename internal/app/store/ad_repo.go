package store

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/emailsender"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
)

type AdRepository interface {
	Create(ad *model.Ad) (model.Ad, error)
	FindByLink(string) (*model.Ad, error)
	CheckPrice()
	UpdatePrices(a *model.Ad, s *emailsender.Sender) error
}
