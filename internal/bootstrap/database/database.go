package database

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
	RD *redis.Client
}

func NewDatabase(DB *sqlx.DB, RD *redis.Client) *Database {
	return &Database{DB, RD}
}