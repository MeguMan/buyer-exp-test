package postgres_store_test

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/MeguMan/buyer-exp-test/internal/app/store"
	"github.com/MeguMan/buyer-exp-test/internal/app/store/postgres_store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := postgres_store.TestDB(t, "user=postgres password=12345 dbname=buyer_exp sslmode=disable")
	defer teardown("users")
	s := postgres_store.New(db)
	u := model.TestUser(t)
	_, err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u.ID)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := postgres_store.TestDB(t, "user=postgres password=12345 dbname=buyer_exp sslmode=disable")
	defer teardown("users")

	s := postgres_store.New(db)
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
