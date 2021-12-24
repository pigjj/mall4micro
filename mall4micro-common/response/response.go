package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	SuccessResponse = &Response{Code: 200, Message: "success"}

	PayloadParseResponse   = &Response{Code: 4000, Message: "payload parse error"}
	UserRegisteredResponse = &Response{Code: 4001, Message: "user already registered"}

	DBConnResponse  = &Response{Code: 6000, Message: "connect database error"}
	SQLExecResponse = &Response{Code: 6010, Message: "sql exec error"}
)
