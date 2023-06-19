package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/wecanooo/gosari/core"
	"strings"
)

func SetupRedis() redis.Cmdable {
	hosts := strings.Split(core.GetConfig().String("REDIS.HOST"), ",")
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    hosts,
		Password: "",
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		fmt.Printf("failed to connect redis: %v\n", err)
		return nil
	}
	fmt.Printf("redis connected successfully: %s\n", hosts)
	return client
}
