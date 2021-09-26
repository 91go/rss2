package gq

import (
	"io"
	"net/http"

	"github.com/gogf/gf/net/ghttp"

	"github.com/91go/rss2/utils"

	"github.com/sirupsen/logrus"

	query "github.com/PuerkitoBio/goquery"
)

const (
	LabelA = "a"
)

// FetchHTML 获取网页
func FetchHTML(url string) *query.Document {
	resp, err := ghttp.NewClient().Get(url)

	if err != nil {
		logrus.WithFields(utils.Fields(url, nil)).Error("http request failed")
		return &query.Document{}
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			logrus.WithFields(utils.Fields(url, nil)).Error("http close failed")
		}
	}(resp.Body)

	return document(url, resp.Response)
}

// PostHTML 发送表单请求
func PostHTML(url string, m map[string]interface{}) *query.Document {
	resp, err := ghttp.NewClient().Post(url, m)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			logrus.WithFields(utils.Fields(url, nil)).Error("http close failed")
		}
	}(resp.Response.Body)

	return document(url, resp.Response)
}

// 请求goquery
func document(url string, resp *http.Response) *query.Document {
	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		logrus.WithFields(utils.Fields(url, err)).Error("goquery failed")
		return &query.Document{}
	}

	return doc
}
