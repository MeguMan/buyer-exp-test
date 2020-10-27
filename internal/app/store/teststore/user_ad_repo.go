package teststore

type UserAdRepository struct {
	store *Store
}

func (r *UserAdRepository) Create(userId, adId int) error {
	return nil
}
