package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		ID:    0,
		Email: "example785123@gmail.com",
	}
}

func TestAd(t *testing.T) *Ad {
	t.Helper()

	return &Ad{
		ID:    0,
		Link:  "example.com",
		Price: 123,
	}
}

func TestUserAd(t *testing.T) *UserAd {
	t.Helper()

	return &UserAd{
		UserId: 0,
		AdId:   0,
	}
}
