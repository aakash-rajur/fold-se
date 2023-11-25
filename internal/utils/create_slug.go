package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateSlug(name string) string {
	hasher := sha256.New()

	hasher.Write([]byte(name))

	buffer := hasher.Sum(nil)

	base16 := hex.EncodeToString(buffer)

	return base16[:7]
}
