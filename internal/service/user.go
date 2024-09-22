package service

import (
	"myapi/internal/bootstrap/database"
	"myapi/internal/repository/users"
	"myapi/internal/request"
	"myapi/internal/response"
)

type UserService struct {
	Db *database.Database
	userRepository users.UserRepository
}

func NewUserService(Db *database.Database, userRepository users.UserRepository) *UserService {
	return &UserService{ Db, userRepository }
}

func (_uc *UserService) UserDetails(in request.UserDetailsIn) (out response.UserDetailsOut, err error) {
	var user users.User
	user, err = _uc.userRepository.FindById(in.UserID)
	if err != nil {
		return
	}

	out.User = &user

	return
}