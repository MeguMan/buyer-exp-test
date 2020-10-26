package postgres_store

type UserAdRepository struct {
	store *Store
}

func (r *UserAdRepository) Create(userId, adId int) error {
	var exists int
	err := r.store.db.QueryRow(
		"SELECT user_id FROM users_ads WHERE user_id = $1 AND ad_id = $2",
		userId, adId,
	).Scan(&exists)

	if exists != 0 {
		return nil
	}

	_, err = r.store.db.Exec("INSERT INTO users_ads (user_id, ad_id) VALUES ($1, $2)",
		userId, adId)

	return err
}
