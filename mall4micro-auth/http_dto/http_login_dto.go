package http_dto

import "github.com/jianghaibo12138/mall4micro/mall4micro-user/dto"

type HttpLoginDTO struct {
	dto.RegisterDTO
	Mobile string `json:"mobile"`
}
