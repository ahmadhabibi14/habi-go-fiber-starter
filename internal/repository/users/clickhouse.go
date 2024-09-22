package users

import (
	"myapi/internal/bootstrap/database"
)

const (
	TABLE_USER = `"users"`
)

type UserImpl_ClickHouse struct {
	Db *database.Database
}

func NewUserImpl_ClickHouse(Db *database.Database) UserImpl_ClickHouse {
	return UserImpl_ClickHouse{Db: Db}
}

func (_u *UserImpl_ClickHouse) Create(user User) (User, error) {
	return User{}, nil
}

func (_u *UserImpl_ClickHouse) UpdateById(id uint64, user User) error {
	return nil
}

func (_u *UserImpl_ClickHouse) FindByEmail(email string) (User, error) {
	return User{}, nil
}

func (_u *UserImpl_ClickHouse) FindById(id uint64) (User, error) {
	return User{}, nil
}

func (_u *UserImpl_ClickHouse) DeleteById(id uint64) error {
	return nil
}