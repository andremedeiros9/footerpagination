package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//function responsible to deal with the request and response of the service
func getPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	query := r.URL.Query()

	currentpage, err := strconv.Atoi(query.Get("currentpage"))
	around, err := strconv.Atoi(query.Get("around"))
	totalpages, err := strconv.Atoi(query.Get("totalpages"))
	boundaries, err := strconv.Atoi(query.Get("boundaries"))
	if err != nil {
		log.Fatal(err)
		return
	}

	footerpagination := newFooter(currentpage, around, totalpages, boundaries)
	json.NewEncoder(w).Encode(footerpagination)
}
