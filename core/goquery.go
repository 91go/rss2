package core

import (
	"net/http"

	"github.com/91go/rss2/utils"

	"github.com/sirupsen/logrus"

	query "github.com/PuerkitoBio/goquery"
)

const (
	LabelA = "a"
)

// CreateClient a http
func CreateClient() *http.Client {
	return &http.Client{}
}

// FetchHTML 获取网页
func FetchHTML(url string) *query.Document {
	client := CreateClient()
	resp, err := client.Get(url)

	if err != nil {
		logrus.WithFields(utils.Fields(url, nil)).Error("http request failed")
		return &query.Document{}
	}

	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		logrus.WithFields(utils.Fields(url, err)).Error("goquery failed")
		return &query.Document{}
	}
	defer resp.Body.Close()
	return doc
}
