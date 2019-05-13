package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}
}
func homePage(w http.ResponseWriter, r *http.Request) { // (1)
	fmt.Fprint(w, "Welcome to the HomePage!") // (2)
}
