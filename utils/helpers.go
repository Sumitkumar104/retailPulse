package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateJobID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
