package domain

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r Response) WithJson(w http.ResponseWriter, LabelPath string) {
	response, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Server is  not ready")); err != nil {
			log.Println("handleFunc/" + LabelPath + "/Marshal/write")
			return
		}
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Println("handleFunc/" + LabelPath + "/write")
		return
	}
}
