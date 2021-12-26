package habit

import (
	"fmt"
	"path/filepath"

	"github.com/gogf/gf/os/gtime"

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
	// todo
	ret = append(ret, rss.Item{
		Title:       fmt.Sprintf("[%s] - %s", gtime.Date(), "生活习惯"),
		Contents:    ReadMarkdown("", "life.md"),
		ID:          rss.GenDateID("habit-md", "habit"),
		UpdatedTime: helper.GetToday(),
	}, rss.Item{
		Title:       fmt.Sprintf("[%s] - %s", gtime.Date(), "一些想法"),
		Contents:    ReadMarkdown("", "thought.md"),
		ID:          rss.GenDateID("habit-md", "thought"),
		UpdatedTime: helper.GetToday(),
	}, rss.Item{
		Title:       fmt.Sprintf("[%s] - %s", gtime.Date(), "thought2"),
		Contents:    ReadMarkdown("", "thought2.md"),
		ID:          rss.GenDateID("habit-md", "thought2"),
		UpdatedTime: helper.GetToday(),
	}, rss.Item{
		Title:       fmt.Sprintf("[%s] - %s", gtime.Date(), "thought2"),
		Contents:    ReadMarkdown("", "thought3.md"),
		ID:          rss.GenDateID("habit-md", "thought2"),
		UpdatedTime: helper.GetToday(),
	})

	return
}

// 读取md
func ReadMarkdown(filename, path string) string {
	abs, err := filepath.Abs(fmt.Sprintf("%s%s%s", "./public/md/", path, filename))
	if err != nil {
		logrus.WithFields(log.Text("", err)).Warn("md file not found")
		return ""
	}
	contents := gfile.GetContents(abs)

	return helper.Md2HTML(contents)
}
