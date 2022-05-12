package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}
