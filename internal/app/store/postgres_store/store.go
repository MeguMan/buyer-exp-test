package postgres_store

import (
	"database/sql"
	"github.com/MeguMan/buyer-exp-test/internal/app/store"
	_ "github.com/lib/pq" // ...
)

type Store struct {
	db *sql.DB
	UserRepository *UserRepository
	AdRepository *AdRepository
	UserAdRepository *UserAdRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository == nil {
		s.UserRepository = &UserRepository{
			store: s,
		}
	}

	return s.UserRepository
}

func (s *Store) Ad() store.AdRepository {
	if s.AdRepository == nil {
		s.AdRepository = &AdRepository{
			store: s,
		}
	}

	return s.AdRepository
}

func (s *Store) UserAd() store.UserAdRepository {
	if s.UserAdRepository == nil {
		s.UserAdRepository = &UserAdRepository{
			store: s,
		}
	}

	return s.UserAdRepository
}