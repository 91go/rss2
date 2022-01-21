package code

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/gogf/gf/os/gtime"

	"github.com/91go/rss2/utils/helper/time"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const (
	GoCnBaseUrl = "https://gocn.vip"
	GoCnUrl     = "https://gocn.vip/topics/"
)

func GoCnRss(ctx *gin.Context) {
	topic := ctx.Param("topic")

	url := fmt.Sprintf("%s%s", GoCnUrl, topic)

	res := rss.Rss(&rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "GopherChina",
			Name:   topic,
		},
		Author:      "gocn",
		UpdatedTime: time.GetToday(),
	}, gocnList(url))

	resp.SendXML(ctx, res)
}

func gocnList(url string) []rss.Item {
	doc := gq.FetchHTML(url)

	wrap := doc.Find(".item-list").Find(".topic")
	var wg sync.WaitGroup

	ret := []rss.Item{}

	wrap.Each(func(i int, sel *query.Selection) {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 爬取失败会直接panic，需要recover起来
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println("panic error.")
				}
			}()

			infos := sel.Find(".infos")
			ta := infos.Find(".title").Find("a")
			title, _ := ta.Attr("title")
			itemUrl, _ := ta.Attr("href")

			// 内容
			itemUrl = fmt.Sprintf("%s%s", GoCnBaseUrl, itemUrl)
			detail := gq.FetchHTML(itemUrl)
			detailHtml, _ := detail.Find(".topic-detail").Html()
			// 处理时间
			timeago := detail.Find(".media-body").Find(".info").Find(".timeago").First().Text()
			re := regexp.MustCompile(`\d+`)
			formatTime := strings.Join(re.FindAllString(timeago, -1), "-")
			strToTimeFormat, _ := gtime.StrToTimeFormat(formatTime, "Y-m-d")

			ret = append(ret, rss.Item{
				Title:       title,
				URL:         itemUrl,
				Contents:    detailHtml,
				UpdatedTime: strToTimeFormat.Time,
			})
		}()
	})

	wg.Wait()

	return ret
}
