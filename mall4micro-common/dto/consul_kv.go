package dto

import (
	"github.com/snownd/cake"
)

type CakeRequestConf struct {
	cake.RequestConfig
	FileName string `param:"file_name"`
	Data     string `body:"application-json"`
}

//
// ConsulKvDTO
// @Description: consul KV交互用DTO
//
type ConsulKvDTO struct {
	LockIndex   int    `json:"LockIndex"`
	Key         string `json:"Key"`
	Flags       int    `json:"Flags"`
	Value       string `json:"Value"`
	CreateIndex int    `json:"CreateIndex"`
	ModifyIndex int    `json:"ModifyIndex"`
}
