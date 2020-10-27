package postgres_store_test

import (
	"database/sql"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/MeguMan/buyer-exp-test/internal/app/store/postgres_store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdRepository_Create(t *testing.T) {
	db, teardown := postgres_store.TestDB(t, "user=postgres password=12345 dbname=buyer_exp sslmode=disable")
	defer teardown("ads")
	s := postgres_store.New(db)
	a := model.TestAd(t)
	_, err := s.Ad().Create(a)
	assert.NoError(t, err)
	assert.NotNil(t, a.ID)
}

func TestAdRepository_FindByLink(t *testing.T) {
	db, teardown := postgres_store.TestDB(t, "user=postgres password=12345 dbname=buyer_exp sslmode=disable")
	defer teardown("ads")

	s := postgres_store.New(db)
	a1 := model.TestAd(t)
	_, err := s.User().FindByEmail(a1.Link)
	assert.EqualError(t, err, sql.ErrNoRows.Error())

	s.Ad().Create(a1)
	a2, err := s.Ad().FindByLink(a1.Link)
	assert.NoError(t, err)
	assert.NotNil(t, a2)
}

func TestAdRepository_UpdatePrices(t *testing.T) {
	db, teardown := postgres_store.TestDB(t, "user=postgres password=12345 dbname=buyer_exp sslmode=disable")
	defer teardown("ads")

	a := model.TestAd(t)
	s := postgres_store.New(db)
	err := s.Ad().UpdatePrices(a)
	assert.NoError(t, err)
}
