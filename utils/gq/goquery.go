package gq

import (
	"io"
	"net/http"
	"rss2/utils/log"

	"github.com/gogf/gf/net/ghttp"

	"github.com/sirupsen/logrus"

	query "github.com/PuerkitoBio/goquery"
)

// FetchHTML 获取网页
// ph2路由需要该方法
func FetchHTML(url string) *query.Document {
	resp, err := http.Get(url)
	if err != nil {
		logrus.WithFields(log.Text(url, nil)).Error("http request failed")
		return &query.Document{}
	}

	defer resp.Body.Close()

	return DocQuery(resp)
}

// FetchHTML 获取网页
func FetchHTMLByGF(url string) *query.Document {
	resp, err := ghttp.NewClient().Get(url)
	if err != nil {
		logrus.WithFields(log.Text(url, nil)).Error("http request failed")
		return &query.Document{}
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			logrus.WithFields(log.Text(url, nil)).Error("http close failed")
		}
	}(resp.Body)

	return DocQuery(resp.Response)
}

// PostHTML 发送表单请求
func PostHTML(url string, m map[string]any) *query.Document {
	resp, err := ghttp.NewClient().Post(url, m)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			logrus.WithFields(log.Text(url, nil)).Error("http close failed")
		}
	}(resp.Response.Body)

	return DocQuery(resp.Response)
}

// 请求goquery
func DocQuery(resp *http.Response) *query.Document {
	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		logrus.WithFields(log.Text(resp.Request.URL.String(), err)).Error("goquery failed")
		return &query.Document{}
	}

	return doc
}
