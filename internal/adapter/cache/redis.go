package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	rds "github.com/redis/go-redis/v9"
	"github.com/yafiesetyo/dating-apps-srv/config"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
)

type redis struct {
	client *rds.Client
}

var _ interfaces.Cache = (*redis)(nil)

func New(cfg config.Config) *redis {
	return &redis{
		client: rds.NewClient(&rds.Options{
			Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
			DB:   cfg.Redis.Db,
		}),
	}
}

func (r *redis) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, rds.Nil) {
		return "", err
	}

	return result, nil
}

func (r *redis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	_, err := r.client.Set(ctx, key, value, ttl).Result()
	return err
}

func (r *redis) Keys(ctx context.Context, pattern string) ([]string, error) {
	var (
		cursor uint64
		result []string
	)

	iter := r.client.Scan(ctx, uint64(cursor), pattern, 0).Iterator()
	for iter.Next(ctx) {
		result = append(result, iter.Val())
	}

	if iter.Err() != nil {
		return result, iter.Err()
	}

	return result, nil
}
