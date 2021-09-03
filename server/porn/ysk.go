package porn

import (
	"fmt"
	"strings"
	"time"

	"github.com/91go/rss2/core"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/text/gregex"
)

const (
	YskURL = "https://yskhd.com/archives/tag/"
)

func YskRss(ctx *gin.Context) {
	tag := ctx.GetString("tag")
	url := fmt.Sprintf("%s%s", YskURL, tag)

	list := parseList(url)

	res := core.Rss(&core.Feed{
		URL:    url,
		Title:  fmt.Sprintf("%s%s", "优丝库-", tag),
		Author: tag,
	}, list)

	ctx.Data(200, "application/xml; charset=utf-8", []byte(res))
}

// 解析列表页
func parseList(url string) []core.Feed {
	doc := core.FetchHTML(url)

	total := doc.Find(".post").Size()
	if total >= core.LimitItem {
		total = core.LimitItem
	}
	wrap := doc.Find(".post").Slice(0, total)
	ret := []core.Feed{}
	wrap.Each(func(i int, selection *query.Selection) {
		href, _ := selection.Find(".img").Find("a").Attr("href")
		title, _ := selection.Find(".img").Find("a").Attr("title")
		cover, _ := selection.Find(".img").Find("a").Find("img").Attr("src")

		ret = append(ret, core.Feed{
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
	cut, _ := gregex.MatchString(".*/(.*)-", url)
	s := cut[1]
	trim := TrimRight(s, s[len(s)-3:])
	parse, err := time.Parse("20060102150405", trim)
	if err != nil {
		return time.Time{}
	}
	return parse
}

// 解析详情页，获取所有图片
func parsePics(url string) string {
	doc := core.FetchHTML(url)
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

func TrimRight(str string, characterMask ...string) string {
	//  characters which are stripped by Trim* functions in default.
	trimChars := string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.TrimRight(str, trimChars)
}
