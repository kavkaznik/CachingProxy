package repository

import (
	"context"
	"fmt"
	"log"

	redis "github.com/redis/go-redis/v9"
)

func Test() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ping)
}
