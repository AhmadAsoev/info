package services

import (
	"info/pkg/domain"
	"net/http"
)

var Infos []domain.Info

func AddInfo(info domain.Info) (response domain.Response) {
	if err := info.Validate(); err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "",
		}
	}

	Infos = append(Infos, info)

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "Success, add information!",
	}
}
