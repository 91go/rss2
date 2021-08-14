package porn

import (
	"fmt"
	"github.com/91go/rss2/utils"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gorilla/feeds"
	"log"
	"strings"
	"time"
)

type YskList struct {
	// 详情url
	url string
	// title
	title string
	// time
	time time.Time
}

var (
	LIMIT = 1
)

func YskRss(request *ghttp.Request) {
	tag := request.GetString("tag")
	url := "https://yskhd.com/archives/tag/" + tag
	list := parseList(url)

	feed := &feeds.Feed{
		Title:   "优丝库——" + tag,
		Link:    &feeds.Link{Href: url},
		Author:  &feeds.Author{Name: tag},
		Created: list[0].time,
	}
	for _, value := range list {

		pics := parsePics(value.url)
		feed.Add(&feeds.Item{
			Title:       value.title,
			Link:        &feeds.Link{Href: value.url},
			Description: pics,
			Author:      &feeds.Author{Name: tag},
			Created:     value.time,
			//Updated:     value.time,
		})
	}

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}

	request.Response.WriteXmlExit(atom)
}

// 解析列表页
func parseList(url string) []YskList {
	doc := utils.FetchHTML(url)

	total := doc.Find(".post").Size()
	if total >= LIMIT {
		total = LIMIT
	}
	wrap := doc.Find(".post").Slice(0, total)
	ret := []YskList{}
	wrap.Each(func(i int, selection *query.Selection) {
		href, _ := selection.Find(".img").Find("a").Attr("href")
		title, _ := selection.Find(".img").Find("a").Attr("title")
		cover, _ := selection.Find(".img").Find("a").Find("img").Attr("src")

		ret = append(ret, YskList{
			url:   href,
			title: title,
			time:  sanitizeTime(cover),
		})
	})

	return ret
}

// 处理时间
func sanitizeTime(url string) time.Time {
	cut, _ := gregex.MatchString(".*/(.*)-", url)
	s := cut[1]
	trim := gstr.TrimRight(s, s[len(s)-3:])
	parse, err := time.Parse("20060102150405", trim)
	if err != nil {
		return time.Time{}
	}
	return parse
}

// 解析详情页，获取所有图片
func parsePics(url string) string {
	doc := utils.FetchHTML(url)
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
