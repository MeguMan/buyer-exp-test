package teststore

type UserAdRepository struct {
	store    *Store
	usersAds map[int]int
}

func (r *UserAdRepository) Create(userId, adId int) error {
	return nil
}
