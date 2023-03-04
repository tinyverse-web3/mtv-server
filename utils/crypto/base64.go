package crypto

import (
	"encoding/base64"
)

func EncryptBase64(data string) string {
	keyB := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(keyB, []byte(data))

	enStr := string(keyB)
	return enStr
}

func DecryptBase64(deData string) string {
	keyB := make([]byte, base64.StdEncoding.DecodedLen(len(deData)))
	n, _ := base64.StdEncoding.Decode(keyB, []byte(deData))

	data := string(keyB[:n])
	return data
}
