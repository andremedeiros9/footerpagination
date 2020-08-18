package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Params struct {
	currentpage int
	around      int
	totalpages  int
	boundaries  int
}

func translateParams(r *http.Request) (params Params, err error) {

	query := r.URL.Query()

	params.currentpage, err = strconv.Atoi(query.Get("currentpage"))
	if err != nil {
		return params, fmt.Errorf("Current Page value is not accepted")
	}
	params.around, err = strconv.Atoi(query.Get("around"))
	if err != nil {
		return params, fmt.Errorf("Around value is not accepted")
	}
	params.totalpages, err = strconv.Atoi(query.Get("totalpages"))
	if err != nil {
		return params, fmt.Errorf("Total Pages value is not accepted")
	}
	params.boundaries, err = strconv.Atoi(query.Get("boundaries"))
	if err != nil {
		return params, fmt.Errorf("Boundaries value is not accepted")
	}

	return params, nil

}

//function responsible to deal with the request and response of the service
func getPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method != http.MethodGet {
		log.Printf("Not a GET method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params, err := translateParams(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("All is well")
	footerpagination := newFooter(params.currentpage, params.around, params.totalpages, params.boundaries)
	json.NewEncoder(w).Encode(footerpagination)
}
