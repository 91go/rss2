package str

import (
	"fmt"
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

// GetIframe 封装iframe标签视频
func GetIframe(videoURL, description string) string {
	return fmt.Sprintf("<iframe src=%s frameborder='0' width='640' height='390' scrolling='no' allowfullscreen></iframe><br><br>%s<br>", videoURL, description)
}
