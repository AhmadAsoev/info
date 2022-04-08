package cmd

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	port := ":8080"

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server is not ready")
	}

}
