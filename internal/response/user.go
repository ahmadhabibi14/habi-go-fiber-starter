package response

import "myapi/internal/repository/users"

type UserDetailsOut struct {
	User *users.User	`json:"user" form:"user"`
} // @name UserDetailsOut