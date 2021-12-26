package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type TokenUtil struct {
	Username string
	Email    string
	Mobile   string
	Status   int
}

var (
	ErrTokenInvalid      = errors.New("token invalid")
	ErrParseToken        = errors.New("parse token")
	ErrUnexpectedSigning = errors.New("unexpected signing method")
)

var hmacSampleSecret = []byte("346b281f-9282-49fe-9970-b18f9fdac659")

func (t *TokenUtil) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": t.Username,
		"email":    t.Email,
		"mobile":   t.Mobile,
		"status":   t.Status,
		"iss":      "mall4micro",
		"iat":      time.Now().Unix(),                                     // iat(issued at): 在什么时候签发的token
		"exp":      time.Now().AddDate(0, 0, 7).Unix(),                    // exp(expires): token什么时候过期
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(), // nbf(not before)：token在此时间之前不能被接收处理
	})
	return token.SignedString(hmacSampleSecret)
}

func (t *TokenUtil) Parse(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigning
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return ErrTokenInvalid
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrParseToken
	}
	t.Username = claims["username"].(string)
	t.Email = claims["email"].(string)
	t.Mobile = claims["mobile"].(string)
	t.Status = int(claims["status"].(float64))
	return nil
}
