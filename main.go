package main

import (
	"c2netapi"
	"log"
	"net/http"
)

func main() {
	router := c2netapi.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
