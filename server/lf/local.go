package lf

import (
	"fmt"
	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gorilla/feeds"
	"strconv"
)

const (
	RootDir = "/srv/"
)

func LocalFileRss(ctx *gin.Context) {
	path := ctx.Param("path")
	host := ctx.Request.Host

	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "local",
			Name:   path,
		},
		UpdatedTime: helper.GetToday(),
	}, FileList(host, path))

	resp.SendXML(ctx, res)
}

func FileList(host, path string) []rss.Item {

	ret := []rss.Item{}
	dstPath := fmt.Sprintf("%s%s", RootDir, path)

	files, err := helper.GetAllFiles(dstPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, file := range files {
		str := gstr.Str(file, RootDir)
		title := gfile.Basename(file)
		fileInfo, err := gfile.Info(file)
		if err != nil {
			return nil
		}
		size := fileInfo.Size()

		tfUrl := fmt.Sprintf("https://%s%s%s", host, RootDir, str)
		filetype := helper.GetContentType(file)
		ret = append(ret, rss.Item{
			Title:    title,
			URL:      tfUrl,
			Contents: DealContents(filetype, tfUrl),
			Enclosure: &feeds.Enclosure{
				Url:    tfUrl,
				Length: strconv.FormatInt(size, 10),
				Type:   filetype,
			},
			UpdatedTime: helper.GetToday(),
		})
	}

	return ret
}

// DealContents 根据文件类型，判断是否返回iframe
func DealContents(filetype, tfUrl string) string {
	if gstr.Contains(filetype, "video") {
		return fmt.Sprintf(`<iframe src="%s" frameborder="0" width="640" height="390" scrolling="no" frameborder="0" border="0" framespacing="0" allowfullscreen></iframe><br><br>`, tfUrl)
	}
	return ""
}
