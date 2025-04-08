package addons

import (
	"crypto/sha256"
	"encoding/hex"
)

func getSHA256Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
func Encrypt(pass string) string {
	hash := getSHA256Hash(pass)
	return hash
}
