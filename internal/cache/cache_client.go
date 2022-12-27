package cache

import (
	"context"

	"github.com/oliver258/eagle/internal/model"
	"github.com/oliver258/eagle/pkg/cache"
	"github.com/oliver258/eagle/pkg/encoding"
	"github.com/oliver258/eagle/pkg/redis"
)

func getCacheClient(ctx context.Context) cache.Cache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	client := cache.NewRedisCache(redis.RedisClient, cachePrefix, jsonEncoding, func() interface{} {
		return &model.UserBaseModel{}
	})

	return client
}
