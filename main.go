package main

import (
	"log"
	"net/http"
)

func main() {
	ports := ":8080"
	router := NewRouter()
	log.Fatal(http.ListenAndServe(ports, router))
}
