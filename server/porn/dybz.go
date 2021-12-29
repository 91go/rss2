package porn

import (
	"fmt"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"

	query "github.com/PuerkitoBio/goquery"

	"github.com/gin-gonic/gin"
)

const (
	DybzBaseUrl   = "http://m.hongrenxs.net"
	DybzBookUrl   = "http://m.hongrenxs.net/book/"
	DybzSearchUrl = "http://m.hongrenxs.net/s.html"
)

// DybzRss 第一版主rss源
func DybzRss(ctx *gin.Context) {
	novel := ctx.Param("novel")
	url := fmt.Sprintf("%s%s/", DybzBookUrl, novel)

	info, list := dybzList(url)
	res := rss.Rss(&info, list)

	resp.SendXML(ctx, res)
}

// 某novel的列表
func dybzList(url string) (feed rss.Feed, feeds []rss.Item) {
	doc := gq.FetchHTML(url)

	wrap := doc.Find(".list_xm").Find("ul").Find("li")
	ret := []rss.Item{}
	wrap.Each(func(i int, selection *query.Selection) {
		title := selection.Find("a").Text()
		novelUrl, _ := selection.Find("a").Attr("href")

		ret = append(ret, rss.Item{
			Title: title,
			URL:   novelUrl,
		})
	})

	info := dybzInfo(url, doc)

	return info, ret
}

func dybzInfo(url string, doc *query.Document) rss.Feed {
	novelName := doc.Find(".cataloginfo").Find("h3").Text()
	author := doc.Find(".infotype").Find("p").Find("a").Text()
	return rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "第一版主",
			Name:   novelName,
		},
		Author: author,
	}
}
