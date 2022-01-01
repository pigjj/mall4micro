package http_dto

import "github.com/pigjj/mall4micro/mall4micro-user/http_dto"

type HttpLoginDTO struct {
	http_dto.RegisterDTO
	Mobile string `json:"mobile"`
}
