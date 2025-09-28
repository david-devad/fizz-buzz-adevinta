package main

import (
	"net/http"

	"example.com/fizz-buzz-adevinta/internal/handler"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/fizzbuzz", handler.FizzBuzzHandler)

    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

	server.ListenAndServe()
}
