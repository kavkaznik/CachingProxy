package controller

import (
	e "main/internal/entities"
	"net/http"
)

type Servicer interface {
	Get(string) (*e.MyRespond, error)
}

type RequestHandler struct {
	Ser Servicer
	URL string
}

func NewRequestHandler(s Servicer, url string) *RequestHandler {
	return &RequestHandler{
		Ser: s,
		URL: url,
	}
}

func (re *RequestHandler) OneHandler(w http.ResponseWriter, r *http.Request) {
	path := "https://" + re.URL + r.URL.Path
	resp, err := re.Ser.Get(path)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Disposition", "inline")
	for key, value := range resp.Headers() {
		w.Header().Set(key, value[0])
	}
	w.Write(resp.Respond())
}
