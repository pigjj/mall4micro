package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	SuccessResponse = &Response{Code: 200, Message: "success"}

	PayloadParseResponse      = &Response{Code: 4000, Message: "payload parse error"}
	UserRegisteredResponse    = &Response{Code: 4001, Message: "user already registered"}
	UserNotRegisteredResponse = &Response{Code: 4002, Message: "user not registered"}
	UserPasswordResponse      = &Response{Code: 4003, Message: "user password incorrect"}
	SignTokenResponse         = &Response{Code: 4004, Message: "sign token error"}

	DBConnResponse  = &Response{Code: 6000, Message: "connect database error"}
	SQLExecResponse = &Response{Code: 6001, Message: "sql exec error"}
)
