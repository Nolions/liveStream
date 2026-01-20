package helper

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// EncryptMD5
// MD5加密
func EncryptMD5(elems []string) string {
	str := strings.Join(elems, "")
	hash := md5.Sum([]byte(str))

	return hex.EncodeToString(hash[:])
}

// EncodeHex
// 將字串透過透過16進位編碼
func EncodeHex(str string) string {
	return hex.EncodeToString([]byte(str))
}
