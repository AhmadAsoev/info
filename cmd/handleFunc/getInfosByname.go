package handleFunc

import (
	"github.com/gorilla/mux"
	"info/pkg/application/services"
	"net/http"
)

func GetInfosByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	response := services.GetInfosByName(name)
	response.WithJson(w, "GetInfosByName")
}
