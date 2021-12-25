package code

import (
	"fmt"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	ss := "编辑于  2021-12-06 19:45:42"
	trim := strings.Trim(ss, "编辑于  ")
	fmt.Println(trim)
}
