package porn

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	url := "https://zzxcx.netzijin.cn/2021/07/20210725144915628-285x285.jpg"
	res := strings.Replace(url, "285x285", "scaled", -1)
	assert.Equal(t, "https://zzxcx.netzijin.cn/2021/07/20210725144915628-scaled.jpg", res)
}

func TestTime(t *testing.T) {
	url := "https://zzxcx.netzijin.cn/2021/06/20210618132842114-scaled.jpg"
	cut, _ := gregex.MatchString(".*/(.*)-", url)
	t.Log(cut)
	s := cut[1]

	// 字符串转time
	local, _ := time.LoadLocation("Local")
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-06-20 18:16:15", local)
	fmt.Println(tt)

	trim := gstr.TrimRight(s, s[len(s)-TimeMillisecondDigit:])
	tt2, _ := time.ParseInLocation("20060102150405", trim, local)
	fmt.Println(tt2)
}

func Test_sanitizeTime(t *testing.T) {
	toTime, err := gtime.StrToTimeFormat("20210430145102", "YmdHis")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(toTime)
}
