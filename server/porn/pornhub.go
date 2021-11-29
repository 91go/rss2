package porn

import (
	"errors"
	"fmt"

	"github.com/91go/rss2/utils/log"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gogf/gf/text/gstr"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// rsshub的pornhub源没有视频地址，无法直接播放，需要跳转才能播放视频，所以重制该feed
func PornhubRss(ctx *gin.Context) {
	model := ctx.Param("model")
	url := fmt.Sprintf("https://rsshub.wrss.top/pornhub/model/%s", model)

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		logrus.WithFields(log.Text("", errors.New("feed parser failed")))
		return
	}

	ret := []rss.Item{}
	for _, item := range feed.Items {
		link := item.Link
		viewKey := gstr.SubStr(link, gstr.Pos(link, "=")+1)
		ret = append(ret, rss.Item{
			Title:    item.Title,
			Contents: fmt.Sprintf(`<iframe src="https://www.pornhub.com/embed/%s" frameborder="0" width="640" height="390" scrolling="no" allowfullscreen></iframe><br><br>%s<br>`, viewKey, item.Description),
			URL:      link,
		})
	}

	res := rss.Rss(&rss.Feed{
		URL:    feed.Link,
		Title:  rss.Title{Prefix: "pornhub", Name: model},
		Author: model,
		Time:   *feed.UpdatedParsed,
	}, ret)

	resp.SendXML(ctx, res)
}
