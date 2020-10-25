package teststore

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/store"
)

type Store struct {
	userRepository   *UserRepository
	adRepository     *AdRepository
	userAdRepository *UserAdRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	return s.userRepository
}

func (s *Store) Ad() store.AdRepository {
	return s.adRepository
}

func (s *Store) UserAd() store.UserAdRepository {
	return s.userAdRepository
}
