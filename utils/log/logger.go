package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	Url = "url"
	Err = "err"
)

func Text(url string, err error) logrus.Fields {
	if err != nil {
		return logrus.Fields{
			"msgtype": "text",
			"text": map[string]string{
				"content": fmt.Sprintf("url: %s \nerror: %s", url, err.Error()),
			},
		}
	}
	return logrus.Fields{
		"msgtype": "text",
		"text": map[string]string{
			"content": fmt.Sprintf("url: %s \nerror: %s", url, ""),
		},
	}
}
