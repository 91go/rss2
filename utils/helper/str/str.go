package str

import (
	"strings"
)

// TrimBlank 移除HTML的空格
func TrimBlank(str string) string {
	t := strings.Replace(str, " ", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "&nbsp", "", -1)
	t = strings.Replace(t, " ", "", -1)
	return t
}
