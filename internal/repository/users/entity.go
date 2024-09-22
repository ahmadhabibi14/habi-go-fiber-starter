package users

import "time"

type User struct {
	Id        uint64    `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}