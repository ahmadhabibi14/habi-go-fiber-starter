package request

type RegisterIn struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
} // @name RegisterIn

type LoginIn struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
} // @name LoginIn

type LogoutIn struct {
	SessionID 	string 			`json:"session_id" form:"session_id" validate:"required"`
} // @name LogoutIn