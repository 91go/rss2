package code

import (
	"fmt"
	"github.com/91go/rss2/core"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"time"
)

//type HuangZ struct {
//	Author string
//	Url    string
//	Title  string
//	Ctx    string
//	Time   time.Time
//}

var (
	url = "https://blog.huangz.me/"
)

// 用来输出rss
// 全文直出会timeout，所以只打出标题
func HuangZRss(ctx *gin.Context) {

	ass := crawlHuangZ()

	res := core.Rss(core.Feed{
		Title: "huangz——黄建宏redis博客",
	}, ass)

	ctx.Data(200, "application/xml; charset=utf-8", []byte(res))
}

// [huangz/blog — blog.huangz.me](https://blog.huangz.me/#)
func crawlHuangZ() []core.Feed {

	doc := core.FetchHTML(url)

	wrap := doc.Find(".toctree-l2")

	var param = []core.Feed{}
	wrap.Each(func(i int, selection *query.Selection) {
		articleUrl, _ := selection.Find(".reference").Attr("href")
		title := selection.Find(".reference").Text()

		fullUrl := fmt.Sprintf("%s%s", url, articleUrl)

		param = append(param, core.Feed{
			Author: "huangz",
			Url:    fullUrl,
			Title:  title,
			Time:   time.Now(),
		})
	})

	return param
}
