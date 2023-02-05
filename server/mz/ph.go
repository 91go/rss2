package mz

import (
	"errors"
	"fmt"
	"rss2/utils/helper/str"
	"rss2/utils/helper/time"
	"rss2/utils/log"
	"rss2/utils/resp"
	"rss2/utils/rss"

	"github.com/gogf/gf/text/gregex"

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
	var num int
	if len(feed.Items) >= DefaultNum {
		num = DefaultNum
	} else {
		num = len(feed.Items)
	}
	for _, item := range feed.Items[0:num] {
		link := item.Link
		viewKey := gstr.SubStr(link, gstr.Pos(link, "=")+1)
		// TODO 优化这里"获取时间"的代码
		updatedTimeArr, _ := gregex.MatchString(`videos\/(.*)\/(original|thumbs_.*)`, item.Description)
		ss, _ := gregex.MatchString(`(.*)\/(.*)\/`, updatedTimeArr[1])
		updatedTime := time.StrToTime(fmt.Sprintf("%s/%s", ss[1], ss[2]), "Ym/d")
		ret = append(ret, rss.Item{
			Title:       item.Title,
			Contents:    str.GetIframe("https://www.pornhub.com/embed/"+viewKey, item.Description),
			URL:         link,
			ID:          viewKey,
			UpdatedTime: updatedTime,
			Author:      model,
		})
	}

	res := rss.Rss(&rss.Feed{
		URL:         feed.Link,
		Title:       rss.Title{Prefix: "pornhub", Name: model},
		Author:      model,
		UpdatedTime: *feed.UpdatedParsed,
	}, ret)

	resp.SendXML(ctx, res)
}
