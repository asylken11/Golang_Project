// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/team", getTeam).Methods("GET")
	router.HandleFunc("/team/player/{name}", getPlayerByName).Methods("GET")
	router.HandleFunc("/health", healthCheck).Methods("GET")

	port := 8080
	log.Printf("Server starting on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Na'Vi App is healthy!")
}
