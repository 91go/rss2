package utils

import (
	"strings"
	"time"

	"github.com/gogf/gf/os/gtime"
)

// GetToday 获取今天的零点时间
func GetToday() time.Time {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t
}

// StrToTime 字符串转time.Time
func StrToTime(str, format string) time.Time {
	st, err := gtime.StrToTimeFormat(str, format)
	if err != nil {
		return time.Time{}
	}
	return st.Time
}

// TrimBlank 移除HTML的空格
func TrimBlank(str string) string {
	t := strings.Replace(str, " ", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "&nbsp", "", -1)
	t = strings.Replace(t, " ", "", -1)
	return t
}
