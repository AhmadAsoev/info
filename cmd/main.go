package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	POST = "POST"
)

func main() {
	router := mux.NewRouter()
	port := ":8080"

	//AddInfo create new rout with POST method
	router.HandleFunc("/info", handleFunc.AddInfo).Methods(POST)

	fmt.Println("server is ready")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server is not ready")
	}

}
