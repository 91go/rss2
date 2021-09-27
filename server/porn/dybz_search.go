package porn

import (
	"errors"
	"fmt"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/log"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/sirupsen/logrus"
)

// DybzSearchRss 搜索某小说
func DybzSearchRss(ctx *gin.Context) {
	novel := ctx.Param("novel")
	m := g.Map{
		"s":    novel,
		"type": "articlename",
	}

	doc := gq.PostHTML(DybzSearchUrl, m)
	url, exists := doc.Find(".searchresult").Find(".sone").Find("a").Attr("href")
	if !exists {
		logrus.WithFields(log.Text(url, errors.New("not exist novel")))
		return
	}

	// 根据id获取最新小说，返回小说url
	info, list := dybzList(fmt.Sprintf("%s%s", DybzBaseUrl, url))
	res := rss.Rss(&info, list)

	resp.SendXML(ctx, res)
}
