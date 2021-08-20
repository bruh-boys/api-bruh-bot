package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1 align="center">Make a http post request</h1>`)
}

func ShowExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phones)
}
