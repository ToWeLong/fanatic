package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

type RedisService interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Delete(key string) error
	Incr(key string) error
	Setnx(key string, value interface{}) error
	Lpush(key string, value interface{}) error
	Lindex(key string, index int64) (string, error)
	Llen(key string) (int64, error)
	Exists(key string) (int64, error)
}

type redisService struct {
}

func (r *redisService) Exists(key string) (int64, error) {
	return RedisClient.Exists(Ctx, key).Result()
}

func (r *redisService) Llen(key string) (int64, error) {
	return RedisClient.LLen(Ctx, key).Result()
}

func (r *redisService) Lindex(key string, index int64) (string, error) {
	return RedisClient.LIndex(Ctx, key, index).Result()
}

func (r *redisService) Lpush(key string, value interface{}) error {
	return RedisClient.LPush(Ctx, key, value).Err()
}

func (r *redisService) Setnx(key string, value interface{}) error {
	return RedisClient.SetNX(Ctx, key, value, 0).Err()
}

func (r *redisService) Delete(key string) error {
	return RedisClient.Del(Ctx, key).Err()
}
func (r *redisService) Incr(key string) error {
	return RedisClient.Incr(Ctx, key).Err()
}

func (r *redisService) Set(key string, value interface{}) error {
	return RedisClient.Set(Ctx, key, value, 0).Err()
}

func (r *redisService) Get(key string) (string, error) {
	return RedisClient.Get(Ctx, key).Result()
}

func NewRedis() RedisService {
	return &redisService{}
}

func Init() {
	if os.Getenv("ENV") == "dev" {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.addr"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		})
	} else {
		db, _ := strconv.Atoi(os.Getenv("REDISDBNAME"))
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDISHOST"),
			Password: os.Getenv("REDISPASSWORD"),
			DB:       db,
		})
	}

	if RedisClient.Ping(Ctx).Err() == nil {
		log.Println("redis 连接成功~")
	}else {
		log.Println("redis 连接失败~")
	}
}

func Close() {
	RedisClient.Close()
}
