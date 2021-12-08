package helper

import (
	"github.com/sirupsen/logrus"
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

func TransTime(str string) time.Time {
	format, err := gtime.StrToTimeFormat(str, "Y/n/d H:i:s")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"time": str,
			"err":  err,
		}).Warn("transTime failed")
		return time.Time{}
	}
	return format.Time
}

// TrimBlank 移除HTML的空格
func TrimBlank(str string) string {
	t := strings.Replace(str, " ", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "&nbsp", "", -1)
	t = strings.Replace(t, " ", "", -1)
	return t
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
