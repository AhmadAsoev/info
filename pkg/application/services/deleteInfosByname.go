package services

import (
	"info/pkg/domain"
	"net/http"
)

func DeleteInfosById(id string) (response domain.Response) {
	var delInfo []domain.Info
	for _, info := range Infos {
		if info.ID.String() != id {
			delInfo = append(delInfo, info)
		}
	}
	Infos = delInfo

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "success deleted",
	}
}
