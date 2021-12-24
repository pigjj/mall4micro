package models

import "time"

//
// MallBase
// @Description: 基础模型
// @Description: 基础模型
//
type MallBase struct {
	ID           uint       `gorm:"column:id;comment:ID;primaryKey;autoIncrement" json:"id"`
	CreateUserId uint       `gorm:"column:create_user_id;comment:创建者ID;default:null" json:"create_user_id"`
	CreateAt     *time.Time `gorm:"column:create_at;comment:创建时间;autoCreateTime" json:"create_at"`
	UpdateAt     *time.Time `gorm:"column:update_at;comment:更新时间;autoCreateTime;autoUpdateTime" json:"update_at"`
	IsDeleted    *uint      `gorm:"column:is_deleted;comment:逻辑更新标志;default:0" json:"is_deleted"`
}

var (
	Deleted = 1
)
