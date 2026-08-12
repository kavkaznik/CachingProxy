package main

import (
	"flag"
	"fmt"
	"log"
	"main/internal/router"
	"net/http"
)

var port int
var url string

func main() {
	flag.IntVar(&port, "port", 8081, "port for url")
	flag.StringVar(&url, "origin", "iana.org", "request address")
	flag.Parse()

	rout := router.NewRouter(url)
	fmt.Printf("Starting proxy\nPort: %v\nOrigin address: %v\nProxy address: localhost:%v\n", port, url, port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%v", port), rout)
	log.Fatal(err)
}
