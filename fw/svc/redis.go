package svc

import "github.com/go-redis/redis/v8"

const Key_Redis = "tk:redis"

type Redis interface {
	// GetClient(opt *redis.Options) (*redis.Client, error)
	GetClient(key string) (*redis.Client, error)
	GetClientShared(key string, idx int) (*redis.Client, error)
}
