package response

import "myapi/internal/repository/users"

type RegisterOut struct {
	User *users.User `json:"user" form:"user"`
} // @name RegisterOut

type LoginOut struct {
	Message string `json:"message" form:"message"`
} // @name LoginOut

type LogoutOut struct {
	Message string `json:"message" form:"message"`
} // @name LogoutOut