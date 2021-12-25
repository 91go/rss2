package code

import (
	"fmt"
	"strings"
	"time"

	"github.com/91go/rss2/utils/helper"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	query "github.com/PuerkitoBio/goquery"

	"github.com/gin-gonic/gin"
)

const (
	NowCoderBaseURL    = "https://ac.nowcoder.com"
	NowCoderDiscussURL = "https://ac.nowcoder.com/discuss/tag/"
)

// https://ac.nowcoder.com/discuss/tag/2656?type=2&order=3
// rsshub不支持tag，只有type和order
func NowCoderRss(ctx *gin.Context) {
	tag := ctx.Param("tag")
	typ := ctx.Param("type")
	order := ctx.Param("order")

	url := fmt.Sprintf("%s%s?type=%s?order=%s", NowCoderDiscussURL, tag, typ, order)

	list := nowCoderList(url)
	res := rss.Rss(&rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "golang",
			Name:   "牛客网",
		},
		UpdatedTime: helper.GetToday(),
	}, list)

	resp.SendXML(ctx, res)
}

func nowCoderList(url string) []rss.Item {
	doc := gq.FetchHTML(url)
	wrap := doc.Find(".common-list").Find(".discuss-detail").Find(".discuss-main").Slice(0, 3)

	ret := []rss.Item{}
	wrap.Each(func(i int, sel *query.Selection) {
		detailURL, _ := sel.Find("a").First().Attr("href")
		title := sel.Find("a").First().Text()

		fullDetailURL := fmt.Sprintf("%s%s", NowCoderBaseURL, detailURL)
		contents, updatedTime := parseDetail(fullDetailURL)
		ret = append(ret, rss.Item{
			Title:       title,
			URL:         fullDetailURL,
			UpdatedTime: updatedTime,
			Contents:    contents,
		})
	})
	return ret
}

func parseDetail(url string) (string, time.Time) {
	doc := gq.FetchHTML(url)
	postTime := doc.Find(".post-time").Text()
	ss := strings.Trim(helper.TrimBlank(postTime), "编辑于")
	toTime := helper.StrToTime(ss, "Y-m-dH:i:s")

	html, _ := doc.Find(".post-topic-main").Find(".post-topic-des").Html()

	return html, toTime
}
