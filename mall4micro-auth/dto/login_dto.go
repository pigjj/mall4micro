package dto

import "github.com/jianghaibo12138/mall4micro/mall4micro-user/dto"

type LoginDTO struct {
	dto.RegisterDTO
	Mobile string `json:"mobile"`
}
