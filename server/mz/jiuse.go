package mz

import (
	"fmt"
	"rss2/utils/gq"
	"rss2/utils/helper/str"
	"rss2/utils/resp"
	"rss2/utils/rss"
	"strings"
	"time"

	"github.com/gogf/gf/text/gstr"

	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
)

const (
	JiuSeBaseURL     = "https://jiuse911.com"
	JiuSeAuthorURL   = "https://jiuse911.com/author/"
	JiuSeKeywordsURL = "https://jiuse911.com/search?keywords="
)

// JiuSeAuthorRss 91porny输出rss
func JiuSeAuthorRss(ctx *gin.Context) {
	author := ctx.Param("author")
	url := fmt.Sprintf("%s%s", JiuSeAuthorURL, author)

	list := jsList(url, author)
	res := rss.Rss(&rss.Feed{
		URL:         url,
		Title:       rss.Title{Prefix: "91porn", Name: author},
		Author:      author,
		UpdatedTime: gtime.Now().Time,
	}, list)

	resp.SendXML(ctx, res)
}

// JiuSeKeywordsRss 91porny输出rss
func JiuSeKeywordsRss(ctx *gin.Context) {
	keywords := ctx.Param("keywords")
	url := fmt.Sprintf("%s%s", JiuSeKeywordsURL, keywords)

	list := jsList(url, keywords)
	res := rss.Rss(&rss.Feed{
		URL:         url,
		Title:       rss.Title{Prefix: "91porn", Name: keywords},
		Author:      keywords,
		UpdatedTime: gtime.Now().Time,
	}, list)

	resp.SendXML(ctx, res)
}

func jsList(url, author string) []rss.Item {
	doc := gq.FetchHTML(url)

	wrap := doc.Find(".colVideoList")

	ret := []rss.Item{}
	wrap.Each(func(i int, selection *query.Selection) {
		title := selection.Find(".video-elem").Find(".title").Text()
		href, _ := selection.Find(".video-elem").Find(".title").Attr("href")
		text := selection.Find(".text-muted").Eq(1).Text()

		videoURL := fmt.Sprintf("%s%s", JiuSeBaseURL, patchVideoURL(href))
		ret = append(ret, rss.Item{
			Title:       title,
			URL:         videoURL,
			Contents:    str.GetIframe(gstr.Replace(videoURL, "view", "embed"), ""),
			UpdatedTime: getCreateTime(text),
			Author:      author,
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
