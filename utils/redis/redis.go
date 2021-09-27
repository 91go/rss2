package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/91go/rss2/utils/config"

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
	fmt.Println("addr===", config.GetString("redis.addr"), "===")
	fmt.Println("pwd===", config.GetString("redis.pwd"), "===")

	conn := redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.Addr"),
		Password: config.GetString("redis.Password"),
		DB:       config.GetInt("redis.DB"),
	})

	if _, err := conn.Ping(Ctx).Result(); err != nil {
		log.Fatalf("connect to redis client failed, err: %v \n", err)
	}

	return conn
}
