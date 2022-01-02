package habit

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/91go/rss2/utils/helper/file"

	"github.com/91go/rss2/utils/helper/html"
	"github.com/91go/rss2/utils/helper/time"

	"github.com/gogf/gf/os/gtime"

	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gfile"
)

var MdFilepath, _ = filepath.Abs("./public/md/")

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
	files, err := file.GetAllFiles(MdFilepath)
	if err != nil {
		fmt.Println("")
		return nil
	}
	for _, filename := range files {
		ret = append(ret, item(filename))
	}
	return
}

func item(fp string) rss.Item {
	filename := strings.TrimPrefix(fp, MdFilepath)

	return rss.Item{
		Title:       fmt.Sprintf("[%s] - %s", gtime.Date(), filename),
		Contents:    ReadMarkdown(fp),
		ID:          rss.GenDateID("habit-md", filename),
		UpdatedTime: time.GetToday(),
	}
}

// 读取md
func ReadMarkdown(fp string) string {
	contents := gfile.GetContents(fp)

	return html.Md2HTML(contents)
}
