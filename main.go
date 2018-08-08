package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	router.NotFoundHandler = http.HandlerFunc(notFoundError)
	log.Fatal(http.ListenAndServe(":8000", router))
}
