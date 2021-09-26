package porn

import (
	"fmt"
	"strings"
	"time"

	"github.com/91go/rss2/core/resp"
	"github.com/91go/rss2/core/rss"

	"github.com/91go/rss2/core/gq"

	"github.com/91go/rss2/utils"
	"github.com/sirupsen/logrus"

	"github.com/gogf/gf/text/gstr"

	"github.com/gogf/gf/os/gtime"

	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/text/gregex"
)

const (
	// 优丝库tag的url
	YskURL    = "https://yskhd.com/archives/tag/"
	TimeDigit = 8
)

// YskRss 优丝库rss [优丝库HD - 提供高清美女写真|丝袜美腿|美女私房|cosplay|美女街拍|4K写真，一站式浏览](https://yskhd.com/)
func YskRss(ctx *gin.Context) {
	tag := ctx.Param("tag")
	url := fmt.Sprintf("%s%s", YskURL, tag)

	list := parseList(url)

	res := rss.Rss(&rss.Feed{
		URL:    url,
		Title:  rss.Title{Prefix: "优丝库", Name: tag},
		Author: tag,
	}, list)

	resp.SendXML(ctx, res)
}

// 解析列表页
func parseList(url string) []rss.Item {
	doc := gq.FetchHTML(url)

	total := doc.Find(".post").Size()
	if total >= rss.LimitItem {
		total = rss.LimitItem
	}
	wrap := doc.Find(".post").Slice(0, total)
	ret := []rss.Item{}
	wrap.Each(func(i int, selection *query.Selection) {
		href, _ := selection.Find(".img").Find(gq.LabelA).Attr("href")
		title, _ := selection.Find(".img").Find(gq.LabelA).Attr("title")
		cover, _ := selection.Find(".img").Find(gq.LabelA).Find("img").Attr("src")

		ret = append(ret, rss.Item{
			URL:      href,
			Title:    title,
			Time:     sanitizeTime(cover),
			Contents: parsePics(href),
		})
	})

	return ret
}

// 处理时间
func sanitizeTime(url string) time.Time {
	cut, err := gregex.MatchString(".*/(.*).", url)
	if err != nil {
		logrus.WithFields(utils.Fields(url, err)).Error("trans time regex failed")
		return time.Time{}
	}
	s := cut[1]
	ts := gstr.TrimRightStr(s, s[TimeDigit:])

	format, err := gtime.StrToTimeFormat(ts, "Ymd")
	if err != nil {
		logrus.WithFields(utils.Fields(url, err)).Error("trans time failed")
		return time.Time{}
	}
	return format.Time
}

// 解析详情页，获取所有图片
func parsePics(url string) string {
	doc := gq.FetchHTML(url)
	wrap := doc.Find(".gallery-fancy-item")
	pics := []string{}
	wrap.Each(func(i int, selection *query.Selection) {
		pic, _ := selection.Find(gq.LabelA).Attr("href")
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
