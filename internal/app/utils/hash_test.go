package utils

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	fmt.Println(BcryptHash("123456"))
}
