package service

import (
	"fmt"
	"io"
	"net/http"
)

type Repositorer interface {
	Get(string) ([]byte, error)
	Set(string, []byte) error
}
type Service struct {
	Repo Repositorer
}

func New(r Repositorer) *Service {
	return &Service{
		Repo: r,
	}
}

//error "redis: nil"

func (s *Service) Get(url string) ([]byte, error) {
	resp, err := s.Repo.Get(url)
	if err != nil {
		if err.Error() == "redis: nil" {
			res, er := httpRequest(url)
			if er != nil {
				return []byte{}, er
			} else {
				err := s.Repo.Set(url, res)
				if err != nil {
					return []byte{}, err
				} else {
					fmt.Println("miss")
					return res, nil
				}
			}
		} else {
			return []byte{}, err
		}
	} else {
		fmt.Println("hit")
		return resp, nil
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
