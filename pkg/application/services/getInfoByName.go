package services

import (
	"info/pkg/domain"
	"net/http"
)

var NewInfos []domain.Info

func GetInfosByName(name string) (response domain.Response) {
	for _, info := range Infos {
		if name == info.Name {
			NewInfos = append(NewInfos, info)
		}
	}
	if NewInfos == nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   "",
			Message: "Card by this name not found",
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: NewInfos,
	}
}
