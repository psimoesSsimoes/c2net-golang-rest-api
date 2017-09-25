package main

import (
	"c2netapi"
	"log"
	"net/http"
)
var db *sql.DB

func main() {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	if err != nil {
		panic()
	}
	router := c2netapi.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
