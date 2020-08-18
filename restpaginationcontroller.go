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

	key := "currentpage"
	value := query.Get(key)
	params.currentpage, err = strconv.Atoi(query.Get("currentpage"))
	if err != nil || params.currentpage < 1 {
		return params, fmt.Errorf("%s %s value is not accepted", key, value, err)
	}

	key = "around"
	value = query.Get(key)
	params.around, err = strconv.Atoi(query.Get("around"))
	if err != nil || params.around < 1 {
		return params, fmt.Errorf("%s %s value is not accepted", key, value, err)
	}

	key = "totalpages"
	value = query.Get(key)
	params.totalpages, err = strconv.Atoi(query.Get("totalpages"))
	if err != nil || params.totalpages < 1 {
		return params, fmt.Errorf("%s %s value is not accepted", key, value, err)
	}

	key = "boundaries"
	value = query.Get(key)
	params.boundaries, err = strconv.Atoi(query.Get("boundaries"))
	if err != nil || params.boundaries < 1 {
		return params, fmt.Errorf("%s %s value is not accepted", key, value, err)
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
