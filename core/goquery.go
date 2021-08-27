package core

import (
	"crypto/tls"
	"fmt"
	query "github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

// CreateClient a http
func CreateClient() *http.Client {
	tr := &http.Transport{
		// 跳过ssl证书检测
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

// 获取网页
func FetchHTML(url string) *query.Document {
	client := CreateClient()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//doc, err := query.NewDocument(url)
	doc, err := query.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalf("Fetch html from %s error: %s", url, err)
	}
	return doc
}

// 获取非utf8的网页
func FetchHTMLNotUtf8(url, charset string) *query.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("fetch html failed")
	}
	defer res.Body.Close()

	//utfBody, err := iconv.NewReader(res.Body, charset, "utf-8")
	//if err != nil {
	//	log.Fatalf("convert %s utf8 error: %s", url, err.Error())
	//}

	doc, err := query.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("fetch url %s error: %s", url, err.Error())
	}

	return doc
}
