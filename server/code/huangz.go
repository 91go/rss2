package code

import (
	"fmt"
	"time"

	"github.com/91go/rss2/core"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const (
	url = "https://blog.huangz.me/"
)

// HuangZRss 用来输出rss
// 全文直出会timeout，所以只打出标题
func HuangZRss(ctx *gin.Context) {
	ass := crawlHuangZ()

	res := core.Rss(&core.Feed{
		Title: "HuangZ-blog",
	}, ass)

	core.SendXML(ctx, res)
}

// [huangz/blog — blog.huangz.me](https://blog.huangz.me/#)
func crawlHuangZ() []core.Feed {
	doc := core.FetchHTML(url)

	wrap := doc.Find(".toctree-l2")

	var param = []core.Feed{}
	wrap.Each(func(i int, selection *query.Selection) {
		articleURL, _ := selection.Find(".reference").Attr("href")
		title := selection.Find(".reference").Text()

		fullURL := fmt.Sprintf("%s%s", url, articleURL)

		param = append(param, core.Feed{
			Author: "huangz",
			URL:    fullURL,
			Title:  title,
			Time:   time.Now(),
		})
	})

	return param
}
