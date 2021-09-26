package utils

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestDingrusHook(t *testing.T) {

	dh, err := NewDingHook(AssembleUrl(), nil)
	if err != nil {
		t.Fatal(err)
	}
	logrus.AddHook(dh)
	logrus.WithFields(logrus.Fields{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": "苍老师结婚啦!",
			"text":  "#### 苍老师结婚啦~~~  \n> 只是发个测试  \n> *测试结束*",
		},
	}).Error()
}
