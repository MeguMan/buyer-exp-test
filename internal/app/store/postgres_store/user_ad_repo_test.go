package postgres_store_test

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/MeguMan/buyer-exp-test/internal/app/store/postgres_store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAdRepository_Create(t *testing.T) {
	db, teardown := postgres_store.TestDB(t, "user=postgres password=12345 dbname=buyer_exp sslmode=disable")
	defer teardown("users_ads", "users", "ads")
	s := postgres_store.New(db)
	ua := model.TestUserAd(t)
	u := model.TestUser(t)
	a := model.TestAd(t)

	ua.UserId, _ = s.User().Create(u)
	ua.AdId, _ = s.Ad().Create(a)

	err := s.UserAd().Create(ua.UserId, ua.AdId)
	assert.NoError(t, err)
	assert.NotNil(t, a.ID)
}
