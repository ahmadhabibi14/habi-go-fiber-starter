package users

type UserRepository interface {
	Create(user User) (User, error)
	UpdateById(id uint64, user User) error
	FindByEmail(email string) (User, error)
	FindById(id uint64)  (User, error)
	DeleteById(id uint64) error
}