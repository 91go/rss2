package main

import (
	"fmt"

	"github.com/91go/rss2/server/habit"

	code2 "github.com/91go/rss2/server/code"

	"github.com/91go/rss2/utils/resp"

	"github.com/sirupsen/logrus"

	"github.com/91go/rss2/middleware"

	life2 "github.com/91go/rss2/server/life"
	lf2 "github.com/91go/rss2/server/localfiles"
	porn2 "github.com/91go/rss2/server/porn"
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

	// 挂载yl文件夹
	r.StaticFS(lf2.RootDir, gin.Dir(lf2.RootDir, true))

	// localfiles路由
	localfiles := r.Group("/localfiles")
	localfiles.GET("/:path", lf2.LocalFileRss)
	localfiles.GET("/:path/:sub", lf2.LocalSecDirFileRss)

	// code路由
	code := r.Group("/code")
	code.GET("/gocn/:topic", code2.GoCnRss)
	code.GET("/nowcoder/discuss/:tag/:type/:order", code2.NowCoderRss)
	code.GET("/onetab/shared/:page", code2.OneTabSharedRSS)
	code.GET("/onetab/txt", code2.OneTabTXTRSS)

	// life路由分组
	life := r.Group("/life")
	life.GET("/weather", life2.WeatherRss)

	// habit
	life.GET("/habit/yearly", habit.HabitYearlyRss)
	life.GET("/habit/md", habit.HabitMDRss)
	life.GET("/habit/everyday", habit.HabitEverydayRss)

	// porn路由分组
	porn := r.Group("/porn")
	porn.GET("/ysk/:tag", porn2.YskRss)
	porn.GET("/dybz/:novel", porn2.DybzRss)
	porn.GET("/pornhub/model/:model", porn2.PornhubRss)

	return r
}
