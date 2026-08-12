package repository

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Repo struct {
	client *redis.Client
}

func (r *Repo) Set(url string, response []byte) error {
	err := r.client.Set(
		ctx,
		url,
		response,
		10*time.Second).Err()
	return err
}

func (r *Repo) Get(url string) ([]byte, error) {
	val, err := r.client.Get(ctx, url).Bytes()
	if err != nil {

		// if errors.Is(err, redis.Nil) "redis: nil"

		return []byte{}, err

	} else {
		return val, nil
	}

}
func New() *Repo {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Repo{
		client: client,
	}
}
