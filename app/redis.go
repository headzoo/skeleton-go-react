package app

import "gopkg.in/redis.v3"

var Redis *redis.Client

func InitRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr: Settings.Redis.Address,
		DB:   Settings.Redis.Database,
	})
	return nil
}
