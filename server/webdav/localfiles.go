package webdav

import (
	"fmt"
	"strconv"

	"github.com/91go/rss2/utils/helper/html"

	"github.com/91go/rss2/utils/helper/file"
	"github.com/91go/rss2/utils/helper/time"

	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gorilla/feeds"
)

const (
	RootDir = "/srv/"
)

func LocalFileRss(ctx *gin.Context) {
	path := ctx.Query("path")
	host := ctx.Request.Host

	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "local",
			Name:   path,
		},
		UpdatedTime: time.GetToday(),
	}, FileList(host, path))

	resp.SendXML(ctx, res)
}

func FileList(host, path string) []rss.Item {
	ret := []rss.Item{}
	dstPath := fmt.Sprintf("%s%s", RootDir, path)

	files, err := file.GetAllFiles(dstPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, filepath := range files {
		str := gstr.Str(filepath, RootDir)
		title := gfile.Basename(filepath)
		fileInfo, err := gfile.Info(filepath)
		if err != nil {
			return nil
		}
		size := fileInfo.Size()

		tfUrl := fmt.Sprintf("https://%s%s%s", host, RootDir, str)
		filetype := file.GetContentType(filepath)
		ret = append(ret, rss.Item{
			Title:    title,
			URL:      tfUrl,
			Contents: html.GetIframe(filetype, tfUrl),
			Enclosure: &feeds.Enclosure{
				Url:    tfUrl,
				Length: strconv.FormatInt(size, 10),
				Type:   filetype,
			},
			UpdatedTime: time.GetToday(),
			// 如果不设置，gorilla会自动设置一个带日期的ID；该rss除非资源位置变更，否则不更新，所以手动设置ID
			ID: fmt.Sprintf("tag:%s", filepath),
		})
	}

	return ret
}
