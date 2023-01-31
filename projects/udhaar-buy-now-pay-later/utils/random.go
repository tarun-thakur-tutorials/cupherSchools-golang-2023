package utils

import (
	"crypto/rand"

	"github.com/spf13/cast"
)

func GenRandom() string {
	randomeCrypto, _ := rand.Prime(rand.Reader, 128)
	return cast.ToString(randomeCrypto)
}
