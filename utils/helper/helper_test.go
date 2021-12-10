package helper

import (
	"fmt"
	"testing"
)

func TestRandStringRunes(t *testing.T) {
	n := 16
	for i := 0; i < 100; i++ {
		runes := RandStringRunes(n)
		fmt.Println(runes)
	}
}
