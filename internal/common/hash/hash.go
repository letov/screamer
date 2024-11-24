package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// Encode рассчет хеша SHA256
func Encode(data *[]byte, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(*data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
