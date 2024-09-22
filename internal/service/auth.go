package service

import (
	"myapi/internal/bootstrap/database"
	"myapi/internal/repository/users"
)

type AuthService struct {
	Db *database.Database
	userRepository users.UserRepository
}

func NewAuthService(Db *database.Database, userRepository users.UserRepository) *AuthService {
	return &AuthService{ Db, userRepository }
}