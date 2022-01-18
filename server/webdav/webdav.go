package webdav

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	file2 "github.com/91go/rss2/utils/helper/file"

	"github.com/91go/rss2/utils/config"
	"github.com/91go/rss2/utils/helper/html"
	"github.com/91go/rss2/utils/helper/time"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"github.com/studio-b12/gowebdav"
)

const (
	WebdavURL = "https://file.wrss.top"
)

// 使用webdav客户端
// 把之前的localfiles也换成webdav，
// 还有阿里云盘，以后还可以挂载其他硬盘
func WebdavRss(ctx *gin.Context) {
	path := ctx.Query("path")

	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "webdav",
			Name:   path,
		},
		UpdatedTime: time.GetToday(),
	}, WebdavList(path))

	resp.SendXML(ctx, res)
}

func WebdavList(path string) []rss.Item {
	webdavURL := config.GetString("WEBDAV.URL")
	user := config.GetString("WEBDAV.USER")
	pwd := config.GetString("WEBDAV.PWD")

	var wg sync.WaitGroup

	ret := []rss.Item{}

	client := gowebdav.NewClient(webdavURL, user, pwd)
	dir, err := client.ReadDir(path)
	if err != nil {
		return ret
	}
	for _, file := range dir {
		wg.Add(1)
		go func(file os.FileInfo) {
			defer wg.Done()

			resourceURL := fmt.Sprintf("%s/%s/%s", WebdavURL, path, file.Name())

			filetype := file2.GetRemoteFileMIMEType(resourceURL, user, pwd)
			ret = append(ret, rss.Item{
				Title:    file.Name(),
				URL:      resourceURL,
				Contents: html.GetIframe(filetype, resourceURL),
				Enclosure: &feeds.Enclosure{
					Url:    resourceURL,
					Length: strconv.FormatInt(file.Size(), 10),
					Type:   filetype,
				},
				UpdatedTime: time.GetToday(),
				ID:          fmt.Sprintf("tag:%s", resourceURL),
			})
		}(file)
	}
	wg.Wait()
	return ret
}
