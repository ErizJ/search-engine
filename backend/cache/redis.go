package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

// InitRedis 初始化 Redis 连接，失败时优雅降级（不影响搜索功能）
func InitRedis(addr, password string) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("⚠️  Redis 不可用，将跳过缓存: %v", err)
		client = nil
		return
	}
	log.Println("Redis 连接成功")
}

// GetClient 返回 Redis 客户端（可能为 nil，调用方需判断）
func GetClient() *redis.Client {
	return client
}

// Set 写入缓存
func Set(key, value string, ttl time.Duration) error {
	if client == nil {
		return nil
	}
	return client.Set(context.Background(), key, value, ttl).Err()
}

// Get 读取缓存，未命中返回 ("", error)
func Get(key string) (string, error) {
	if client == nil {
		return "", redis.Nil
	}
	return client.Get(context.Background(), key).Result()
}

// Del 删除缓存键
func Del(keys ...string) error {
	if client == nil {
		return nil
	}
	return client.Del(context.Background(), keys...).Err()
}
