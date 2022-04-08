package services

import (
	"info/pkg/domain"
	"net/http"
)

func GetAllInfos() (response domain.Response) {
	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: Infos,
	}
}
