package main

import (
	"flag"
	"fmt"
	"net/http"

	"joaor.dev.com/joao/stringservice/pkg/api"
)

func main() {
	port := flag.Int("port", 12345, "the port for the server")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.HandleParse)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
}
