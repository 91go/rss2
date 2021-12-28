package code

import (
	"fmt"
	"github.com/91go/rss2/utils/helper"
	"sync"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const (
	GoCnBaseUrl = "https://gocn.vip"
	GoCnUrl     = "https://gocn.vip/topics/"
)

func GoCnRss(ctx *gin.Context) {
	topic := ctx.Param("topic")

	url := fmt.Sprintf("%s%s", GoCnUrl, topic)

	res := rss.Rss(&rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "gocn",
			Name:   "title",
		},
		Author:      "gocn",
		UpdatedTime: helper.GetToday(),
	}, gocnList(url))

	resp.SendXML(ctx, res)
}

func gocnList(url string) []rss.Item {
	doc := gq.FetchHTML(url)

	wrap := doc.Find(".item-list").Find(".topic")
	var wg sync.WaitGroup

	ret := []rss.Item{}

	wrap.Each(func(i int, sel *query.Selection) {
		wg.Add(1)

		go func() {
			defer wg.Done()

			ta := sel.Find(".infos").Find(".title").Find("a")
			title, _ := ta.Attr("title")
			itemUrl, _ := ta.Attr("href")

			// 内容
			itemUrl = fmt.Sprintf("%s%s", GoCnBaseUrl, itemUrl)
			detail := gq.FetchHTML(itemUrl)
			detailHtml, _ := detail.Find(".topic-detail").Html()

			ret = append(ret, rss.Item{
				Title:       title,
				URL:         itemUrl,
				Contents:    detailHtml,
				UpdatedTime: helper.GetToday(),
			})
		}()

	})

	wg.Wait()

	return ret
}

// func goCnInfo(url string, doc *query.Document) rss.Feed {
// 	title := doc.Find(".sub-navbar").Find(".container").Find(".summary").Find("p").Text()
// 	return rss.Feed{
//
// 	}
// }
