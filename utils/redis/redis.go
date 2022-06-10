package redis

import (
	"context"
	"fmt"
	"log"
	"rss2/utils/config"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	Conn *redis.Client
}

var Ctx = context.Background()

func NewClient(conn *redis.Client) *Client {
	return &Client{
		Conn: conn,
	}
}

func Conn() *redis.Client {
	redisAddr := config.GetString("REDIS.ADDR")
	redisPWD := config.GetString("REDIS.PASSWORD")
	redisDB := config.GetInt("REDIS.DB")

	fmt.Println("addr===", redisAddr, "===")
	fmt.Println("pwd===", redisPWD, "===")

	conn := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPWD,
		DB:       redisDB,
	})

	if _, err := conn.Ping(Ctx).Result(); err != nil {
		log.Fatalf("connect to redis client failed, err: %v \n", err)
	}

	return conn
}
