package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis(ip, port, password string, db int) *redis.Client {
	// 1. 初始化客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     ip + ":" + port, // Redis 地址
		Password: password,        // 密码
		DB:       db,              // 默认数据库
	})

	// 2. 测试连接 (Ping)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect redis:" + err.Error())
	}

	return rdb
}
