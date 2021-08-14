package life

import (
	"fmt"
	"github.com/91go/rss2/utils"
	"testing"
)

// 字符串转time.Time
func TestTime(t *testing.T) {

	tt := utils.TransTime("2017-06-20 18:16:15")
	fmt.Println(tt)
}
