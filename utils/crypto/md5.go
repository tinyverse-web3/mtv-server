package crypto

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	data := []byte(str)
	b := md5.Sum(data)
	pass := fmt.Sprintf("%x", b)
	return pass
}
