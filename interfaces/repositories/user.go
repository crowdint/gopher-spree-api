package repositories

type UserRepository DbRepo

func NewUserRepository() *UserRepository {
	return &UserRepository{spree_db}
}
