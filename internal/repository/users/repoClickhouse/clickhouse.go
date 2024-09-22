package repoclickhouse

import (
	"myapi/internal/bootstrap/database"
	"myapi/internal/repository/users"
)

const (
	TABLE_USER = `"users"`
)

type UserImpl struct {
	Db *database.Database
}

func NewUserImpl(Db *database.Database) UserImpl {
	return UserImpl{Db: Db}
}

func (_u *UserImpl) Create(user users.User) (users.User, error) {
	return users.User{}, nil
}

func (_u *UserImpl) UpdateById(id uint64, user users.User) error {
	return nil
}

func (_u *UserImpl) FindByEmail(email string) (users.User, error) {
	return users.User{}, nil
}

func (_u *UserImpl) FindById(id uint64) (users.User, error) {
	return users.User{}, nil
}

func (_u *UserImpl) DeleteById(id uint64) error {
	return nil
}