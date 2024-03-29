package mz

import (
	"fmt"
	"rss2/utils/gq"
	"rss2/utils/helper/str"
	time2 "rss2/utils/helper/time"
	"rss2/utils/resp"
	"rss2/utils/rss"
	"time"

	query "github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/text/gregex"

	"github.com/gin-gonic/gin"
)

const PornhubHomepage = "https://cn.pornhub.com"

// rsshub的pornhub源没有视频地址，无法直接播放，需要跳转才能播放视频，所以重制该feed
func PornhubRss(ctx *gin.Context) {
	model := ctx.Param("model")
	url := fmt.Sprintf(PornhubHomepage+"/model/%s/videos?o=mr", model)

	list := parseModelList(url)

	res := rss.Rss(&rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "pornhub",
			Name:   model,
		},
		Author:      model,
		UpdatedTime: time2.GetToday(),
	}, list)

	resp.SendXML(ctx, res)
}

// 解析列表页
func parseModelList(url string) []rss.Item {
	doc := gq.FetchHTML(url)

	wrap := doc.Find("#mostRecentVideosSection .videoBox")
	ret := []rss.Item{}
	wrap.Each(func(i int, selection *query.Selection) {
		vkey, _ := selection.Attr("data-video-vkey")
		base := selection.Find("span.title a")
		href, _ := base.Attr("href")
		title := base.Text()
		des, _ := selection.Find("div").Find(".phimage").Find("a").Find("img").Html()
		dateStr, _ := selection.Find("div").Find(".phimage").Find("a").Find("img").Attr("src")

		ret = append(ret, rss.Item{
			URL:         PornhubHomepage + href,
			Title:       title,
			Contents:    str.GetIframe("https://www.pornhub.com/embed/"+vkey, des),
			UpdatedTime: getTime(dateStr),
			ID:          rss.GenFixedID("ph", vkey),
		})
	})

	return ret
}

func getTime(dateStr string) time.Time {
	updatedTimeArr, _ := gregex.MatchString(`videos\/(.*)\/(original|thumbs_.*)`, dateStr)
	ss, _ := gregex.MatchString(`(.*)\/(.*)\/`, updatedTimeArr[1])
	updatedTime := time2.StrToTime(fmt.Sprintf("%s/%s", ss[1], ss[2]), "Ym/d")
	return updatedTime
}
