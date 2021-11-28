package gq

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/91go/rss2/utils/log"

	"github.com/gogf/gf/net/ghttp"

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
		logrus.WithFields(log.Text(url, nil)).Error("http request failed")
		return &query.Document{}
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			logrus.WithFields(log.Text(url, nil)).Error("http close failed")
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
			logrus.WithFields(log.Text(url, nil)).Error("http close failed")
		}
	}(resp.Response.Body)

	return document(url, resp.Response)
}

func RestyFetchHTML(url string) *query.Document {
	client := resty.New()
	resp, err := client.R().EnableTrace().Get("https://cn.pornhub.com/model/mai-chen/videos?o=mr")
	if err != nil {
		return nil
	}
	fmt.Println(resp)
	fmt.Println(resp.RawResponse.Body)

	if err != nil {
		logrus.WithFields(log.Text(url, nil)).Error("http request failed")
		return &query.Document{}
	}

	return document(url, resp.RawResponse)
}

// 请求goquery
func document(url string, resp *http.Response) *query.Document {
	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		logrus.WithFields(log.Text(url, err)).Error("goquery failed")
		return &query.Document{}
	}

	return doc
}
