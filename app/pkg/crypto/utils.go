package crypto

import (
	"crypto/rand"
	"math/big"
)

func GetRandomValues(size uint) ([]byte, error) {
	values := make([]byte, 0, size)
	for i := 0; i < cap(values); i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(1<<8))
		if err != nil {
			return values, err
		}
		values = append(values, byte(num.Uint64()))
	}
	return values, nil
}
