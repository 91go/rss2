package porn

import (
	"fmt"
	"strings"
	"time"

	"github.com/91go/rss2/core/resp"
	"github.com/91go/rss2/core/rss"

	"github.com/91go/rss2/core/gq"

	"github.com/91go/rss2/utils"

	"github.com/gogf/gf/os/gtime"

	query "github.com/PuerkitoBio/goquery"

	"github.com/91go/gofc"

	"github.com/gin-gonic/gin"
)

const (
	JiuSeBaseURL   = "https://jiuse911.com"
	JiuSeAuthorURL = "https://jiuse911.com/author/"
)

// JiuSeRss 91porny输出rss
func JiuSeRss(ctx *gin.Context) {
	author := ctx.Param("author")
	url := fmt.Sprintf("%s%s", JiuSeAuthorURL, author)

	list := jsList(url)
	res := rss.Rss(&rss.Feed{
		URL:    url,
		Title:  rss.Title{Prefix: "91porn", Name: author},
		Author: author,
		Time:   gtime.Now().Time,
	}, list)

	resp.SendXML(ctx, res)
}

func jsList(url string) []rss.Item {
	doc := gq.FetchHTML(url)

	total := doc.Find(".colVideoList").Size()
	size := gofc.If(total >= rss.LimitItem, rss.LimitItem, total).(int)
	wrap := doc.Find(".colVideoList").Slice(0, size)
	ret := []rss.Item{}
	wrap.Each(func(i int, selection *query.Selection) {
		title := selection.Find(".video-elem").Find(".title").Text()
		href, _ := selection.Find(".video-elem").Find(".title").Attr("href")
		text := selection.Find(".text-muted").Eq(1).Text()

		ret = append(ret, rss.Item{
			Title: title,
			URL:   fmt.Sprintf("%s%s", JiuSeBaseURL, patchVideoURL(href)),
			Time:  getCreateTime(text),
		})
	})
	return ret
}

func patchVideoURL(url string) string {
	contains := strings.Contains(url, "view")
	if contains {
		return strings.Replace(url, "viewhd", "view", -1)
	}
	return url
}

func getCreateTime(text string) time.Time {
	s := strings.Split(text, "|")
	s2 := s[0]

	format, err := gtime.StrToTimeFormat(utils.TrimBlank(s2), "Y-m-d")
	if err != nil {
		return time.Time{}
	}
	return format.Time
}
