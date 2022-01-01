package ctx

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pigjj/mall4micro/mall4micro-common/log"
	"github.com/pigjj/mall4micro/mall4micro-common/response"
	"github.com/pigjj/mall4micro/mall4micro-user/http_dto"
	"net/http"
)

type GinContext struct {
	*gin.Context
	Logger *log.ZapLogger
	User   *http_dto.UserDTO
}

const UserInfoKey = "user"

//
// NewGinContext
// @Description: 构建自定义context
// @param fn
// @param logger
// @return gin.HandlerFunc
//
func NewGinContext(fn func(gtx *GinContext), logger *log.ZapLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(&GinContext{ctx, logger, nil})
	}
}

func SetUser(c *gin.Context, user http_dto.UserDTO) {
	c.Set(UserInfoKey, user)
}

//
// GetUser
// @Description: 从context中获取user信息
// @Document:
// @receiver gtx
// @return *http_dto.UserDTO
// @return error
//
func (gtx *GinContext) GetUser() (*http_dto.UserDTO, error) {
	userItl, ok := gtx.Get(UserInfoKey)
	if !ok {
		return nil, errors.New(response.UserNotLoginResponse.Message)
	}
	user, ok := userItl.(http_dto.UserDTO)
	if !ok {
		return nil, errors.New(response.UserNotLoginResponse.Message)
	}
	return &user, nil
}

//
// Json
// @Description: 返回json数据，不携带额外数据
// @receiver gtx
// @param res
// s
func (gtx *GinContext) Json(res *response.Response) {
	gtx.Context.JSON(http.StatusOK, gin.H{"code": res.Code, "message": res.Message})
}

//
// JsonWithData
// @Description: 返回json数据，携带数据
// @receiver gtx
// @param res
// @param data
//
func (gtx *GinContext) JsonWithData(res *response.Response, data interface{}) {
	gtx.Context.JSON(http.StatusOK, gin.H{"code": res.Code, "message": res.Message, "data": data})
}
