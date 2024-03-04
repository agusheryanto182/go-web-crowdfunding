package redis

import (
	"context"

	"time"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/config"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/caching"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type redisCacheRepository struct {
	rdb *redis.Client
}

func NewRedisClient(rdb *redis.Client) caching.CacheRepository {
	return &redisCacheRepository{
		rdb: rdb,
	}
}

func InitialRedis(cnf config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.Redis.Addr,
		Password: cnf.Redis.Pass,
		DB:       0,
	})

	// Tes koneksi ke Redis
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatalf("Error koneksi Redis: %s", err.Error())
	}

	logrus.Info("Redis connected successfully")

	return rdb
}

func (r redisCacheRepository) Get(key string) ([]byte, error) {
	val, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (r redisCacheRepository) Set(key string, entry []byte, expiration time.Duration) error {
	return r.rdb.Set(context.Background(), key, entry, expiration).Err()
}
