package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// InitRedis 初始化连接redis
func InitRedis()  {
	rdb:=redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	if _,err:=rdb.Ping(ctx).Result();err!=nil {
		fmt.Println("[InitRedis] init redis failed")
		return
	}
	fmt.Println("[InitRedis] init redis success")
}
