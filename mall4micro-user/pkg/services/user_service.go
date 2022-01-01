package services

import (
	"github.com/pigjj/mall4micro/mall4micro-common/conn"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-common/utils"
	"github.com/pigjj/mall4micro/mall4micro-user/dao/mall_sys_user"
	"github.com/pigjj/mall4micro/mall4micro-user/http_dto"
	"gorm.io/gorm"
)

//
// RegisterSrv
// @Description: 用户注册服务
// @Document:
// @param r
// @return *response.Response
// @return error
//
func RegisterSrv(r *http_dto.RegisterDTO) (*response.Response, error) {
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
