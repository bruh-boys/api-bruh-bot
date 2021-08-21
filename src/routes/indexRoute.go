package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1 align="center">make a http post request, more info in <a href="https://github.com/bruh-boys/api-bruh-bot">Github Repository</a></h1>`)
}

func ShowExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phones)
}

func EmailShowExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}
