package service

import (
	"io"
	"net/http"

	e "main/internal/entities"
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

// error "redis: nil"

func (s *Service) Get(url string) (*e.MyRespond, error) {
	resp, err := s.Repo.Get(url)
	if err != nil {
		if err.Error() == "redis: nil" {
			res, er := httpRequest(url)
			if er != nil {
				return res, er
			} else {
				bytes, err := res.Encode()
				if err != nil {
					return res, err
				}
				err = s.Repo.Set(url, bytes)
				if err != nil {
					return res, err
				} else {
					res.Miss()
					return res, nil
				}
			}
		} else {
			return nil, err
		}
	} else {
		mr := e.NewMyRespond()
		err = mr.Decode(resp)
		mr.Hit()
		return mr, err
	}
}

func httpRequest(url string) (*e.MyRespond, error) {
	mr := e.NewMyRespond()
	resp, err := http.Get(url)
	if err != nil {
		return mr, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return mr, err
	}
	mr.Body = body
	mr.Header = resp.Header
	return mr, err
}
