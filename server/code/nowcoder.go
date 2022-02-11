package code

import (
	"fmt"
	"strings"
	"time"

	"github.com/91go/rss2/utils/helper/str"
	time2 "github.com/91go/rss2/utils/helper/time"

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

var (
	tagMap = map[string]string{
		"2656": "golang",
	}
	typeMap = map[string]string{
		"0": "全部",
		"2": "笔经面经",
	}
	orderMap = map[string]string{
		"0": "最新回复",
		"3": "最新发表",
		"1": "最热",
		"4": "精华",
	}
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
			Prefix: "牛客网",
			Name:   fmt.Sprintf("%s/%s/%s", tagMap[tag], typeMap[typ], orderMap[order]),
		},
		UpdatedTime: time2.GetToday(),
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
			ID:          rss.GenFixedID("nowcoder", fullDetailURL),
		})
	})
	return ret
}

func parseDetail(url string) (string, time.Time) {
	doc := gq.FetchHTML(url)
	postTime := doc.Find(".post-time").Text()
	html, _ := doc.Find(".post-topic-main").Find(".post-topic-des").Html()

	if strings.Contains(postTime, "编辑于") {
		ss := strings.Trim(str.TrimBlank(postTime), "编辑于")
		toTime := time2.StrToTime(ss, "Y-m-dH:i:s")

		return html, toTime
	}

	return html, time2.GetToday()
}
