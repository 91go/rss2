package porn

import (
	"fmt"
	"strings"

	"github.com/91go/rss2/utils/gq"
	time2 "github.com/91go/rss2/utils/helper/time"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"

	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const (
	// 优丝库tag的url
	YskURL = "https://yskhd.com/archives/tag/"
)

// YskRss 优丝库rss [优丝库HD - 提供高清美女写真|丝袜美腿|美女私房|cosplay|美女街拍|4K写真，一站式浏览](https://yskhd.com/)
func YskRss(ctx *gin.Context) {
	tag := ctx.Param("tag")
	url := fmt.Sprintf("%s%s", YskURL, tag)

	list := parseList(url)

	res := rss.Rss(&rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "优丝库",
			Name:   tag,
		},
		Author:      tag,
		UpdatedTime: time2.GetToday(),
	}, list)

	resp.SendXML(ctx, res)
}

// 解析列表页
func parseList(url string) []rss.Item {
	doc := gq.FetchHTML(url)

	wrap := doc.Find(".post")
	ret := []rss.Item{}
	wrap.Each(func(i int, selection *query.Selection) {
		href, _ := selection.Find(".img").Find("a").Attr("href")
		title, _ := selection.Find(".img").Find("a").Attr("title")

		ret = append(ret, rss.Item{
			URL:         href,
			Title:       title,
			Contents:    parsePics(href),
			UpdatedTime: time2.GetToday(),
			ID:          rss.GenFixedID("ysk", href),
		})
	})

	return ret
}

// 处理时间
// func sanitizeTime(url string) time.Time {
// 	cut, err := gregex.MatchString(".*/(.*).", url)
// 	if err != nil {
// 		logrus.WithFields(log.Text(url, err)).Error("trans time regex failed")
// 		return time.Time{}
// 	}
// 	s := cut[1]
// 	ts := gstr.TrimRightStr(s, s[TimeDigit:])
//
// 	format, err := gtime.StrToTimeFormat(ts, "Ymd")
// 	if err != nil {
// 		logrus.WithFields(log.Text(url, err)).Error("trans time failed")
// 		return time.Time{}
// 	}
// 	return format.Time
// }

// 解析详情页，获取所有图片
func parsePics(url string) string {
	doc := gq.FetchHTML(url)
	wrap := doc.Find(".gallery-fancy-item")
	pics := []string{}
	wrap.Each(func(i int, selection *query.Selection) {
		pic, _ := selection.Find("a").Attr("href")
		pics = append(pics, pic)
	})

	wrap2 := doc.Find(".gallery-blur-item")
	wrap2.Each(func(i int, selection *query.Selection) {
		origPic, _ := selection.Find("span").Find("img").Attr("src")
		// 替换为scaled
		pic := strings.Replace(origPic, "285x285", "scaled", -1)
		pics = append(pics, pic)
	})

	ret := ""
	for _, pic := range pics {
		ret += fmt.Sprintf("<img src=%s>", pic)
	}
	return ret
}
