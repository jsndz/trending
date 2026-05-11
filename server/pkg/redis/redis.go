package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return rdb
}

func InvalidateTrendingCache(ctx context.Context, rdb *redis.Client) error {

	iter := rdb.Scan(ctx, 0, "trending:*", 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()

		if err := rdb.Del(ctx, key).Err(); err != nil {
			return err
		}
	}

	return iter.Err()
}
