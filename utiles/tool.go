package utiles

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"time"
)

// Encryption 密码加密
func Encryption(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
// CreateToken 生成token
func CreateToken(pwd string) (string, error) {
	encrpty := []byte(time.Now().String() + pwd)
	h := sha256.New()
	_, err := h.Write(encrpty)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

