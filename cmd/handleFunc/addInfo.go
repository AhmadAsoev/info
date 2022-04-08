package handleFunc

import (
	"encoding/json"
	"info/pkg/application/services"
	"info/pkg/domain"
	"log"
	"net/http"
)

func AddInfo(w http.ResponseWriter, r *http.Request) {
	var request domain.Info

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Can not decode into json")); err != nil {
			log.Println("handleFunc/AddInfo/Decode/write")
			return
		}
	}
	response := services.AddInfo(request)
	response.WithJson(w, "AddInfo")
}
