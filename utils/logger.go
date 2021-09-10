package utils

import "github.com/sirupsen/logrus"

const (
	Url = "url"
	Err = "err"
)

func Fields(url string, err error) logrus.Fields {
	return logrus.Fields{
		"url": url,
		"err": err,
	}
}
