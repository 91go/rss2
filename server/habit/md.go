package habit

import (
	"fmt"
	"path/filepath"

	"github.com/91go/rss2/utils/helper/html"
	"github.com/91go/rss2/utils/helper/time"

	"github.com/gogf/gf/os/gtime"

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
		UpdatedTime: time.GetToday(),
	}, DietFeed())

	resp.SendXML(ctx, res)
}

func DietFeed() (ret []rss.Item) {
	ret = append(ret, item("life"), item("thought"), item("thought2"), item("thought3"))
	return
}

func item(title string) rss.Item {
	return rss.Item{
		Title:       fmt.Sprintf("[%s] - %s", gtime.Date(), title),
		Contents:    ReadMarkdown(fmt.Sprintf("%s.md", title)),
		ID:          rss.GenDateID("habit-md", title),
		UpdatedTime: time.GetToday(),
	}
}

// 读取md
func ReadMarkdown(path string) string {
	abs, err := filepath.Abs(fmt.Sprintf("%s%s", "./public/md/", path))
	if err != nil {
		logrus.WithFields(log.Text("", err)).Warn("md file not found")
		return ""
	}
	contents := gfile.GetContents(abs)

	return html.Md2HTML(contents)
}
