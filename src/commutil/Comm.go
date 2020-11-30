package commutil

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// 获取md5
func GetMd5(data *[]byte) string {
	m := md5.Sum(*data)
	return hex.EncodeToString(m[:])
}

// 生成base64
func EncodeBase64(data *[]byte) string {
	return base64.StdEncoding.EncodeToString(*data)
}

// 解析base64
func DecodeBase64(text string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(text)
}
