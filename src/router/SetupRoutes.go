package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bruh-boys/api-bruh-bot/src/controller"
	"github.com/gorilla/mux"
)

func SetupRoutes() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", controller.IndexRoute)
	r.HandleFunc("/sms", controller.ShowExample).Methods("GET")
	r.HandleFunc("/sms", controller.SmsSpam).Methods("POST")
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}
	fmt.Printf("Api on port: %s", port)
	http.ListenAndServe(":"+port, r)
}
