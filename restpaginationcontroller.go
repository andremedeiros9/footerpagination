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

	if r.Method != http.MethodGet {
		log.Printf("Not a GET method")
		w.WriteStatus(http.BadRequest)
		return
	}

	query := r.URL.Query()

	currentpage, err := strconv.Atoi(query.Get("currentpage"))
	if err != nil {
		w.WriteStatus(http.StatusBadRequest)
		log.Printf("Current Page value not accepted")
		return
	}
	around, err := strconv.Atoi(query.Get("around"))
	if err != nil {
		w.WriteStatus(http.StatusBadRequest)
		log.Printf("Around value not accepted")
		//log.Fatal(err)
		return
	}
	totalpages, err := strconv.Atoi(query.Get("totalpages"))
	if err != nil {
		w.WriteStatus(http.StatusBadRequest)
		log.Printf("Total Pages value not accepted")
		//log.Fatal(err)
		return
	}
	boundaries, err := strconv.Atoi(query.Get("boundaries"))
	if err != nil {
		w.WriteStatus(http.StatusBadRequest)
		log.Printf("Boundaries value not accepted")
		//log.Fatal(err)
		return
	}
	w.WriteStatus(http.StatusOK)
	log.Printf("All is well")
	footerpagination := newFooter(currentpage, around, totalpages, boundaries)
	json.NewEncoder(w).Encode(footerpagination)
}
