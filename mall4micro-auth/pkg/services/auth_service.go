package services

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/http_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conn"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/utils"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/dao/mall_sys_user"
)

func LoginService(d *http_dto.HttpLoginDTO) (string, *response.Response, error) {
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

func AuthorizedService(token string) (*response.Response, *utils.TokenUtil, error) {
	var t utils.TokenUtil
	err := t.Parse(token)
	if err != nil {
		return response.ParseTokenResponse, nil, err
	}
	return response.SuccessResponse, &t, err
}
