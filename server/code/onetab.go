package code

import (
	"fmt"
	"rss2/utils/gq"
	"rss2/utils/helper/time"
	"rss2/utils/resp"
	"rss2/utils/rss"
	"sync"

	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const OneTabBaseURL = "https://www.one-tab.com/page/"

func OneTabSharedRSS(ctx *gin.Context) {
	page := ctx.Param("page")
	url := fmt.Sprintf("%s%s", OneTabBaseURL, page)

	res := rss.Rss(&rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "onetab",
			Name:   page,
		},
		UpdatedTime: time.GetToday(),
	}, sharedList(url))

	resp.SendXML(ctx, res)
}

func sharedList(url string) []rss.Item {
	doc := gq.FetchHTML(url)
	ret := []rss.Item{}

	var wg sync.WaitGroup

	doc.Find("body").Find("div").Slice(7, -1).Each(func(i int, sel *query.Selection) {
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println("panic error.")
				}
			}()

			title := sel.Find("a").Text()
			if url, exists := sel.Find("a").Attr("href"); exists {
				// 获取文章内容
				detail, _ := gq.FetchHTML(url).Find("body").Html()
				ret = append(ret, rss.Item{
					Title:       title,
					URL:         url,
					ID:          rss.GenFixedID("onetab-shared", url),
					Contents:    detail,
					UpdatedTime: time.GetToday(),
				})
			}
		}()
	})

	wg.Wait()

	return ret
}
