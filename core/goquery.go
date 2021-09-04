package core

import (
	"fmt"
	"log"
	"net/http"

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
		fmt.Println(err)
		return &query.Document{}
	}

	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Fetch html from %s error: %s", url, err)
	}
	defer resp.Body.Close()
	return doc
}
