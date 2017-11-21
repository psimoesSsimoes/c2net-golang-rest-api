package c2netapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.PathPrefix("/site/").Handler(http.StripPrefix("/site/", http.FileServer(http.Dir("/home/pi/Github/c2net-golang-rest-api/c2net-iot-hub-site/"))))

	return router
}
