package fake

import (
	"fmt"
	"testing"
)

func TestGetUserAgent(t *testing.T) {
	for i := 0; i < 10; i++ {
		ua := GetUserAgent()
		fmt.Println(ua)
	}
}
