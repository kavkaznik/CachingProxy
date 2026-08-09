package main

import (
	"flag"
	"io"
	"log"
	"main/internal/repository"
	"net/http"
)

var port int
var url string

func main() {
	repository.Test()
	return
	flag.IntVar(&port, "port", 0, "port for url")
	flag.StringVar(&url, "origin", "localhost", "request address")
	flag.Parse()

	http.HandleFunc("/", mo)

	err := http.ListenAndServe("localhost:8081", nil)
	log.Fatal(err)
}

func mo(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(url + r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}
