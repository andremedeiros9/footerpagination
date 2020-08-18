package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getPages).Queries("currentpage", "{currentpage}", "around", "{around}", "totalpages", "{totalpages}", "boundaries", "{boundaries}").Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
