package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type GenCode interface {
	GenCode(length int) (string, error)
}

type genCodeService struct {
}

// NewGenCode creates a new gencode service
func NewGenCode() GenCode {
	return &genCodeService{}
}

// GenCode generate a new code
func (gen *genCodeService) GenCode(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	code := make([]byte, length)
	maxLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, maxLength)

		if err != nil {
			return "", err
		}

		code[i] = charset[n.Int64()]
	}

	return string(code), nil
}
