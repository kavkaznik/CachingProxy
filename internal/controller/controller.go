package controller

import "net/http"

type Servicer interface {
	Get(string) ([]byte, error)
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
	path := re.URL + r.URL.Path
	resp, err := re.Ser.Get(path)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
