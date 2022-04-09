package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"info/cmd/handleFunc"
	"log"
	"net/http"
)

const (
	POST   = "POST"
	GET    = "GET"
	DELETE = "DELETE"
)

func main() {
	router := mux.NewRouter()
	port := ":8080"

	//AddInfo create new rout with POST method
	router.HandleFunc("/info", handleFunc.AddInfo).Methods(POST)

	//GetAllInfos create new rout with GET method
	router.HandleFunc("/allInfos", handleFunc.GetAllInfo).Methods(GET)

	//GetInfosByName create new rout with Get method
	router.HandleFunc("/infosBy/{name}", handleFunc.GetInfosByName).Methods(GET)

	fmt.Println("server is ready")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server is not ready")
	}

}
