package life

import (
	"fmt"
	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/log"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gfile"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func DietRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "吃饭(饮食方案)",
		},
		Author:      "lry",
		UpdatedTime: helper.GetToday(),
	}, DietFeed())

	resp.SendXML(ctx, res)
}

func DietFeed() (ret []rss.Item) {

	ret = append(ret, rss.Item{
		Title:    "diet",
		Contents: ReadMarkdown("diet.md"),
		Time:     helper.GetToday(),
	})

	return
}

func ReadMarkdown(filename string) string {
	abs, err := filepath.Abs(fmt.Sprintf("%s%s", "./public/md/", filename))
	if err != nil {
		logrus.WithFields(log.Text("", err)).Warn("md file not found")
		return ""
	}
	contents := gfile.GetContents(abs)

	return helper.Md2HTML(contents)
}
