package bootstrap

import (
	"fmt"
	"goWeb/pkg/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", "localhost", "6379"),
		"",
		"",
		0,
	)
}
