package request

type UserDetailsIn struct {
	UserID uint64 `json:"user_id" form:"user_id" validate:"required"`
} // @name UserDetailsIn