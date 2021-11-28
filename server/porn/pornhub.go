package porn

import (
	"fmt"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
)

// rsshub的pornhub源没有视频地址，无法直接播放，需要跳转才能播放视频
func PornhubRss(ctx *gin.Context) {
	model := ctx.Param("model")
	url := fmt.Sprintf("https://cn.pornhub.com/model/%s/videos?o=mr", model)

	list := pornhubList(url)
	res := rss.Rss(&rss.Feed{
		URL:    url,
		Title:  rss.Title{Prefix: "pornhub", Name: ""},
		Author: model,
		Time:   gtime.Now().Time,
	}, list)

	resp.SendXML(ctx, res)
}

func pornhubList(url string) []rss.Item {
	doc := gq.RestyFetchHTML(url)
	// wrap := doc.Find("#mostRecentVideosSection").Find(".pcVideoListItem")
	wrap := doc.Find("#mostRecentVideosSection").Find(".videoBox")
	ret := []rss.Item{}

	wrap.Each(func(i int, selection *query.Selection) {
		on := selection.Find(".wrap").Find(".thumbnail-info-wrapper").Find("span")
		title, _ := on.Attr("title")
		videoUrl, _ := on.Attr("href")
		videoFrame := fmt.Sprintf(`<iframe src="https://www.pornhub.com/embed/%s" frameborder="0" width="640" height="390" scrolling="no" allowfullscreen></iframe>`, videoUrl)

		ret = append(ret, rss.Item{
			Title:    title,
			URL:      videoUrl,
			Contents: videoFrame,
		})
	})
	return ret
}