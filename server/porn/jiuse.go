package porn

import (
	"fmt"
	"github.com/gogf/gf/text/gstr"
	"strings"
	"time"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/helper/str"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
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
		URL:         url,
		Title:       rss.Title{Prefix: "91porn", Name: author},
		Author:      author,
		UpdatedTime: gtime.Now().Time,
	}, list)

	resp.SendXML(ctx, res)
}

func jsList(url string) []rss.Item {
	doc := gq.FetchHTML(url)

	wrap := doc.Find(".colVideoList")

	ret := []rss.Item{}
	wrap.Each(func(i int, selection *query.Selection) {
		title := selection.Find(".video-elem").Find(".title").Text()
		href, _ := selection.Find(".video-elem").Find(".title").Attr("href")
		text := selection.Find(".text-muted").Eq(1).Text()

		ret = append(ret, rss.Item{
			Title:       title,
			URL:         fmt.Sprintf("%s%s", JiuSeBaseURL, patchVideoURL(href)),
			Contents:    fmt.Sprintf("<iframe src=%s frameborder='0' width='640' height='340' scrolling='no' allowfullscreen></iframe>", gstr.Replace(url, "view", "embed")),
			UpdatedTime: getCreateTime(text),
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

	format, err := gtime.StrToTimeFormat(str.TrimBlank(s2), "Y-m-d")
	if err != nil {
		return time.Time{}
	}
	return format.Time
}
