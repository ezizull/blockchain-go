package database

import (
	"blockchain-go/infrastructure/repository/redis"
)

type Database struct {
	Redis *redis.InfoDatabaseRedis
}
