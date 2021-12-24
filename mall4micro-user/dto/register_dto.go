package dto

import "errors"

type RegisterDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Platform uint   `json:"platform"`
}

var (
	ErrUsername = errors.New("username invalid")
	ErrPassword = errors.New("password invalid")
)

func (r *RegisterDTO) UsernameValidate() error {
	if len(r.Username) < 5 || len(r.Username) >= 20 {
		return ErrUsername
	}
	return nil
}

func (r *RegisterDTO) PasswordValidate() error {
	if len(r.Password) < 5 || len(r.Username) >= 20 {
		return ErrPassword
	}
	return nil
}
