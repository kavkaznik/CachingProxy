package router

import (
	"net/http"

	"main/internal/controller"
	"main/internal/repository"
	"main/internal/service"

	"github.com/go-chi/chi"
)

func NewRouter(url string, clear bool) http.Handler {
	rep := repository.New(clear)
	ser := service.New(rep)
	con := controller.NewRequestHandler(ser, url)
	r := chi.NewRouter()
	r.HandleFunc("/*", con.OneHandler)

	return r
}
