package repository

import (
	"context"
	"errors"
	"io"
	"net/http"
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
		10*time.Minute).Err()
	return err
}

func (r *Repo) Get(url string) ([]byte, error, string) {
	val, err := r.client.Get(ctx, url).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			body, err := httpRequest(url)
			if err != nil {
				return []byte{}, err, ""
			} else {
				return body, nil, ""
			}
		} else {
			return []byte{}, err, ""
		}
	} else {
		return val, nil, "cache"
	}

}
func NewRepo() *Repo {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Repo{
		client: client,
	}
}

func httpRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body, err
}
