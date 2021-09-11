package main

import (
	"fmt"
	"net/http"
	"os"

	rl "github.com/ahmedash95/ratelimit"
	"github.com/bruh-boys/api-bruh-bot/src/routes"
	"github.com/gorilla/mux"
)

var ratelimit rl.Limit

func main() {
	// Create the limiter
	ratelimit = rl.CreateLimit("1r/m")
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", routes.IndexRoute)
	r.HandleFunc("/sms", routes.ShowExample).Methods("GET")
	r.HandleFunc("/sms", routes.SmsSpam).Methods("POST")
	r.HandleFunc("/email", routes.EmailShowExample).Methods("GET")
	r.HandleFunc("/email", routes.SpamEmail).Methods("POST")
	port, ok := os.LookupEnv("PORT")

	if ok == false {
		port = "5000"
	}
	fmt.Printf("Api on port: %s", port)
	http.ListenAndServe(":"+port, r)
}
