package main

import (
	"github.com/91go/rss2/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

// goleak
// func TestMain(m *testing.M) {
// 	goleak.VerifyTestMain(m)
// }

// http://127.0.0.1:8090/porn/dybz/85867
//
// http://127.0.0.1:8090/porn/dybz/85984
//
// http://127.0.0.1:8090/porn/dybz/search/%E9%80%8D%E9%81%A5%E5%B0%8F%E6%95%A3%E4%BB%99
//
// http://127.0.0.1:8090/life/weather?city=shanghai/shanghai
//
// http://127.0.0.1:8090/porn/ysk/xiuren%E7%A7%80%E4%BA%BA%E7%BD%91
func TestDingTalkHook(t *testing.T) {

	dh, err := utils.NewDingHook(utils.AssembleUrl(), nil)
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
