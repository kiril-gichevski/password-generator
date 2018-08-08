package main

import (
	"./web"
	"log"
	"net/http"
)

func main() {
	router := web.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(web.NotFoundError)
	log.Fatal(http.ListenAndServe(":8000", router))
}
