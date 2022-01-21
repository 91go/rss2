package gq

import (
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/91go/rss2/utils/log"

	"github.com/sirupsen/logrus"

	query "github.com/PuerkitoBio/goquery"
)

// FetchHTML 获取网页
func FetchHTML(url string) *query.Document {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		logrus.WithFields(log.Text(url, nil)).Error("http request failed")
		return &query.Document{}
	}
	if err != nil {
		logrus.WithFields(log.Text(url, nil)).Error("http request failed")
		return &query.Document{}
	}

	return DocQuery(resp.RawResponse)
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
