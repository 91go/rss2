package main

import (
	"fmt"

	code2 "github.com/91go/rss2/server/code"

	"github.com/91go/rss2/utils/resp"

	"github.com/sirupsen/logrus"

	"github.com/91go/rss2/middleware"

	life2 "github.com/91go/rss2/server/life"
	porn2 "github.com/91go/rss2/server/mz"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	gin.SetMode(gin.DebugMode)

	if err := r.Run(":8090"); err != nil {
		fmt.Printf("startup service failed, err:%v \n", err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())

	r.GET("/ping", func(ctx *gin.Context) {
		logrus.WithFields(logrus.Fields{
			"msgtype": "markdown",
			"markdown": map[string]string{
				"title": "ping",
				"text":  "pong",
			},
		}).Error()

		resp.SendJSON(ctx, "pong")
	})

	// code路由
	code := r.Group("/code")
	code.GET("/gocn/:grade", code2.GoCnRss)
	code.GET("/nowcoder/discuss/:tag/:type/:order", code2.NowCoderRss)
	code.GET("/onetab/shared/:page", code2.OneTabSharedRSS)
	code.GET("/codersway", code2.CodersWayRes)

	// life路由分组
	life := r.Group("/life")
	life.GET("/weather", life2.WeatherRss)

	// porn路由分组
	porn := r.Group("/mz")
	porn.GET("/ysk/:tag", porn2.YskRss)
	porn.GET("/dybz/:novel", porn2.DybzRss)
	porn.GET("/pornhub/model/:model", porn2.PornhubRss)
	porn.GET("/jiuse/:author", porn2.JiuSeRss)

	return r
}
