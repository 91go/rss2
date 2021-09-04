package life

import (
	"fmt"
	"testing"
	"time"
)

// 字符串转time.Time
func TestTime(t *testing.T) {
	tt := TransTime3("2021/8/17 18:25:21")
	fmt.Println(tt)
}

func TransTime3(str string) time.Time {
	tt, _ := time.Parse("2006/1/02 15:04:05", str)
	return tt
}
