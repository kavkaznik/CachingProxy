package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"main/internal/router"
)

var (
	port int
	url  string
)

func main() {
	flag.IntVar(&port, "port", 8081, "port for url")
	flag.StringVar(&url, "origin", "iana.org", "request address")
	clear := flag.Bool("clear", false, "clears cache")
	flag.Parse()
	if *clear {
		fmt.Println("Cache is cleared")
	}
	rout := router.NewRouter(url, *clear)
	fmt.Printf("Starting proxy\nPort: %v\nOrigin address: %v\nProxy address: localhost:%v\n", port, url, port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%v", port), rout)
	log.Fatal(err)
}
