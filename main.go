package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	// r.HandleFunc("/webhook", webhookGetHandler).Methods("GET")
	// r.HandleFunc("/webhook", webhookPostHandler).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error listening and server:", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprint(w, "Server Up")
}
