package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	//r := mux.NewRouter()

	//r.HandleFunc("/", getPages).Queries("currentpage", "{currentpage}", "around", "{around}", "totalpages", "{totalpages}", "boundaries", "{boundaries}").Methods("GET")

	authHandler := httpauth.SimpleBasicAuth("Footer", "footer")

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(getPages)
	mux.Handle("/", authHandler(finalHandler))

	//http.HandleFunc("/", authHandler(getPages))

	log.Fatal(http.ListenAndServe(":8000", mux))
}
