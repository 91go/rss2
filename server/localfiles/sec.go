package localfiles

import (
	"fmt"

	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
)

func LocalSecDirFileRss(ctx *gin.Context) {
	path := fmt.Sprintf("%s/%s", ctx.Param("path"), ctx.Param("sub"))
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
