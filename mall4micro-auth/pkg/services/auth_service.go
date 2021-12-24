package services

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/utils"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_user"
)

func LoginService(d *dto.LoginDTO) (string, *response.Response, error) {
	session, err := conn.Conn()
	if err != nil {
		return "", response.DBConnResponse, err
	}
	var user mall_sys_user.MallSysUser
	err = user.SelectUserByUsername(session, d.Username)
	if err != nil {
		return "", response.UserNotRegisteredResponse, err
	}
	var p utils.PasswordUtil
	if ok, err := p.Equal(d.Password, user.Password, user.SaltStr); !ok {
		return "", response.UserPasswordResponse, err
	}
	var tu = utils.TokenUtil{
		Username: user.Username,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Status:   *user.Status,
	}
	token, err := tu.Generate()
	if err != nil {
		return "", response.SignTokenResponse, err
	}
	return token, response.SuccessResponse, nil
}
