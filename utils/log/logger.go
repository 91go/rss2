package log

import "github.com/sirupsen/logrus"

const (
	Url = "url"
	Err = "err"
)

func Text(url string, err error) logrus.Fields {
	return logrus.Fields{
		"msgtype": "text",
		"text":    map[string]interface{}{"url": url, "err": err},
	}
}
