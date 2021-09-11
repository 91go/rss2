package porn

import (
	"fmt"
	"strings"
	"time"

	"github.com/91go/rss2/utils"

	"github.com/gogf/gf/os/gtime"

	query "github.com/PuerkitoBio/goquery"

	"github.com/91go/gofc"

	"github.com/91go/rss2/core"

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
	res := core.Rss(&core.Feed{
		URL:    url,
		Title:  fmt.Sprintf("%s%s", "91porn-", author),
		Author: author,
	}, list)

	core.SendXML(ctx, res)
}

func jsList(url string) []core.Feed {
	doc := core.FetchHTML(url)

	total := doc.Find(".colVideoList").Size()
	size := gofc.If(total >= core.LimitItem, core.LimitItem, total).(int)
	wrap := doc.Find(".colVideoList").Slice(0, size)
	ret := []core.Feed{}
	wrap.Each(func(i int, selection *query.Selection) {
		title := selection.Find(".video-elem").Find(".title").Text()
		href, _ := selection.Find(".video-elem").Find(".title").Attr("href")
		text := selection.Find(".text-muted").Eq(1).Text()

		ret = append(ret, core.Feed{
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
