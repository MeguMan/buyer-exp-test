package store

type UserAdRepository interface {
	Create(userId, adId int) error
}

