package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(plain string) string {
	h := md5.New()
	h.Write([]byte(plain))
	return hex.EncodeToString(h.Sum(nil))
}
