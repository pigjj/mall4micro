package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/google/uuid"
)

var alternativeStr = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

//
// PasswordUtil
// @Description: 密码相关结构体
//
type PasswordUtil struct {
}

//
// Generate
// @Description: 密码加密函数
// @receiver p
// @param password
//
func (p *PasswordUtil) Generate(password string) (string, string) {
	s := uuid.New().String()
	return p.generate(password, s), s
}

//
// generate
// @Description: 生成hash过的密码
// @receiver p
// @param password
// @param saltStr
// @return string
//
func (p *PasswordUtil) generate(password, saltStr string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	m := md5.New()
	saltStr = hex.EncodeToString(bs) + saltStr
	m.Write([]byte(saltStr))
	return hex.EncodeToString(m.Sum(nil))
}

//
// Equal
// @Description: 判断输入密码与DB密码是否相等
// @receiver p
// @param password
// @param hashed
// @param saltStr
// @return bool
//
func (p *PasswordUtil) Equal(password, hashed, saltStr string) bool {
	return hashed == p.generate(password, saltStr)
}
