package services

import (
	"github.com/pigjj/mall4micro/mall4micro-auth/http_dto"
	"github.com/pigjj/mall4micro/mall4micro-common/ctx"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-common/utils"
)

func LoginService(d *http_dto.HttpLoginDTO, gtx *ctx.GinContext) (string, *response.Response, error) {
	if d == nil {
		return "", response.PayloadParseResponse, nil
	}
	user, err := userInfoByUsernameService(d.Username, gtx.Logger)
	if err != nil {
		return "", response.RPCExecResponse, err
	}

	var p utils.PasswordUtil
	if ok, err := p.Equal(d.Password, user.Password, user.SaltStr); !ok {
		return "", response.UserPasswordResponse, err
	}
	var tu = utils.TokenUtil{
		ID:       uint(user.Id),
		Username: user.Username,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Status:   int(user.Status),
	}
	token, err := tu.Generate()
	if err != nil {
		return "", response.SignTokenResponse, err
	}
	gtx.Logger.Infof("[LoginService] user %s login.", user.Username)
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
