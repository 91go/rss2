package life

import (
	"fmt"
	"path/filepath"

	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/log"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gfile"
	"github.com/sirupsen/logrus"
)

func HabitMDRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "生活习惯md",
		},
		Author:      "lry",
		UpdatedTime: helper.GetToday(),
	}, DietFeed())

	resp.SendXML(ctx, res)
}

func DietFeed() (ret []rss.Item) {
	ret = append(ret, rss.Item{
		Title:       "生活习惯",
		Contents:    ReadMarkdown("life.md"),
		UpdatedTime: helper.GetToday(),
	}, rss.Item{
		Title:       "吃饭",
		Contents:    ReadMarkdown("diet.md"),
		UpdatedTime: helper.GetToday(),
	}, rss.Item{
		Title:       "运动",
		Contents:    ReadMarkdown("exercise.md"),
		UpdatedTime: helper.GetToday(),
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
