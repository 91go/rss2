package habit

import (
	"fmt"
	"testing"
)

func TestCheckDateTime(t *testing.T) {
	time := CheckDateTime("2h")
	fmt.Println(time.String())
}
