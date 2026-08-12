package router

import (
	"net/http"

	"main/internal/controller"
	"main/internal/repository"
	"main/internal/service"

	"github.com/go-chi/chi"
)

func NewRouter(url string) http.Handler {
	rep := repository.New()
	ser := service.New(rep)
	con := controller.NewRequestHandler(ser, url)
	r := chi.NewRouter()
	r.HandleFunc("/*", con.OneHandler)

	return r
}
