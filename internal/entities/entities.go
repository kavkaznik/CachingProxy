package entities

import (
	"encoding/json"
)

type MyRespond struct {
	Header map[string][]string `json:"header"`
	Body   []byte              `json:"body"`
}

func NewMyRespond() *MyRespond {
	return &MyRespond{
		Header: make(map[string][]string),
		Body:   make([]byte, 0),
	}
}

func (mr *MyRespond) Headers() map[string][]string {
	return mr.Header
}

func (mr *MyRespond) Respond() []byte {
	return mr.Body
}

func (mr *MyRespond) Decode(body []byte) error {
	return json.Unmarshal(body, mr)
}

func (mr *MyRespond) Encode() ([]byte, error) {
	return json.Marshal(mr)
}

func (mr *MyRespond) Hit() {
	mr.Header["X-Cache"] = []string{"HIT"}
}

func (mr *MyRespond) Miss() {
	mr.Header["X-Cache"] = []string{"MISS"}
}
