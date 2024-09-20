package md5

import (
	"crypto/md5"
	"encoding/hex"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Md5V
//@description: md5加密
//@param: str []byte
//@return: string

func Md5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
