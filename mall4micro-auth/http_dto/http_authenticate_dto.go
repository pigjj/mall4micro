package http_dto

type HttpAuthenticateDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Status   int    `json:"status"`
	SaltStr  string `json:"salt_str"`
}
