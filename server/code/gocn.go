package code

import (
	"fmt"

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

	feed, list := goCnList(url)
	res := rss.Rss(&feed, list)

	resp.SendXML(ctx, res)
}

func goCnList(url string) (feed rss.Feed, list []rss.Item) {
	doc := gq.FetchHTML(url)
	wrap := doc.Find(".item-list").Find(".topic").Slice(0, 3)
	ret := []rss.Item{}
	wrap.Each(func(i int, sel *query.Selection) {
		ta := sel.Find(".infos").Find(".title").Find("a")
		title, _ := ta.Attr("title")
		itemUrl, _ := ta.Attr("href")

		// 内容
		itemUrl = fmt.Sprintf("%s%s", GoCnBaseUrl, itemUrl)
		detail := gq.FetchHTML(itemUrl)
		detailHtml, _ := detail.Find(".topic-detail").Html()

		ret = append(ret, rss.Item{
			Title:    title,
			URL:      itemUrl,
			Contents: detailHtml,
		})
	})

	info := goCnInfo(url, doc)
	return info, ret
}

func goCnInfo(url string, doc *query.Document) rss.Feed {
	title := doc.Find(".sub-navbar").Find(".container").Find(".summary").Find("p").Text()
	return rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "gocn",
			Name:   title,
		},
		Author: "gocn",
	}
}
