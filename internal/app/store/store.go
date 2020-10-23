package store

type Store interface {
	User() UserRepository
	Ad() AdRepository
	UserAd() UserAdRepository
}