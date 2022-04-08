package handleFunc

import (
	"info/pkg/application/services"
	"net/http"
)

func GetAllInfo(w http.ResponseWriter, r *http.Request) {

	response := services.GetAllInfos()
	response.WithJson(w, "GetAllInfo")
}
