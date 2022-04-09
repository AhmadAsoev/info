package handleFunc

import (
	"github.com/gorilla/mux"
	"info/pkg/application/services"
	"net/http"
)

func DeleteInfosById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response := services.DeleteInfosById(id)
	response.WithJson(w, "DeleteInfosById")
}
