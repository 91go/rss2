package code

import (
	"fmt"
	"rss2/utils/helper/time"
	"rss2/utils/req"
	"rss2/utils/resp"
	"rss2/utils/rss"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/container/gmap"
	"github.com/tidwall/gjson"
)

const (
	GoCnBaseURL        = "https://gocn.vip/topics?grade="
	GoCNIndexAPI       = "https://gocn.vip/apiv3/topic/list?currentPage=1&cate2Id=0&grade="
	GoCNArticleAPI     = "https://gocn.vip/apiv3/topic/%s/info"
	GoCNArticleBaseURL = "https://gocn.vip/topics/"
)

var GradeMap = map[string]string{
	"excellent": "精华",
	"hot":       "最热",
	"new":       "最新",
}

// FIXME GoCN改版了，需要重写

// GoCnRss GoCN RSS
func GoCnRss(ctx *gin.Context) {
	grade := ctx.Param("grade")
	if !gmap.NewStrStrMapFrom(GradeMap).Contains(grade) {
		return
	}
	url := fmt.Sprintf("%s%s", GoCNIndexAPI, grade)

	res := rss.Rss(&rss.Feed{
		URL: fmt.Sprintf("%s%s", GoCnBaseURL, grade),
		Title: rss.Title{
			Prefix: "GoCN",
			Name:   GradeMap[grade],
		},
		Author:      "GoCN",
		UpdatedTime: time.GetToday(),
	}, articleList(url))

	resp.SendXML(ctx, res)
}

func articleList(url string) []rss.Item {
	res, _ := req.Get(url)
	lists := gjson.Get(res, "data.list").Array()
	var wg sync.WaitGroup
	ret := []rss.Item{}

	for _, list := range lists {
		wg.Add(1)

		go func(list gjson.Result) {
			defer wg.Done()

			guid := list.Get("guid").String()
			title := list.Get("title").String()
			detailRes, _ := req.Get(fmt.Sprintf(GoCNArticleAPI, guid))
			zw := gjson.Get(detailRes, "data.topic.contentHtml").String()
			articleURL := fmt.Sprintf("%s%s", GoCNArticleBaseURL, guid)
			ret = append(ret, rss.Item{
				Title:    title,
				URL:      articleURL,
				Contents: zw,
				ID:       rss.GenFixedID("GoCN", articleURL),
			})
		}(list)
	}
	wg.Wait()
	return ret
}
