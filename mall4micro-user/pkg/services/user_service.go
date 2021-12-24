package services

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/utils"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_user"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dto"
	"gorm.io/gorm"
)

func RegisterSrv(r *dto.RegisterDTO) (*response.Response, error) {
	session, err := conn.Conn()
	if err != nil {
		return response.DBConnResponse, err
	}
	var user mall_sys_user.MallSysUser
	err = user.SelectUserByUsername(session, r.Username)
	if err == nil {
		return response.UserRegisteredResponse, err
	}
	var p utils.PasswordUtil
	user.Username = r.Username
	user.Password, user.SaltStr = p.Generate(r.Password)

	user.Status = &mall_sys_user.StatusDisable
	err = session.Transaction(func(tx *gorm.DB) error {
		return user.Insert(tx)
	})
	if err != nil {
		return response.SQLExecResponse, err
	}
	return response.SuccessResponse, nil
}
