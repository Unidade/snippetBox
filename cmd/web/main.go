package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("GET /snippet/view", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe("127.0.0.1:4000", mux)
	log.Fatal(err)
}
