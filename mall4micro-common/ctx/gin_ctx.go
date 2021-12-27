package ctx

import (
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/mall4micro/mall4micro-auth/http_dto"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/log"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"net/http"
)

type GinContext struct {
	*gin.Context
	Logger *log.ZapLogger
	User   *http_dto.HttpAuthenticateDTO
}

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
