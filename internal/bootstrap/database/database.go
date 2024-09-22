package database

import (
	"myapi/configs"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
	RD *redis.Client
}

func NewDatabase() (db *Database, err error) {
	var chDb *sqlx.DB
	chDb, err = configs.ConnectClickhouse()
	if err != nil {
		return
	}

	rd := configs.NewRedisClient()
	_, err = rd.Ping().Result()
	if err != nil {
		return	
	}

	db = &Database{
		DB: chDb, RD: rd,
	}

	return
}