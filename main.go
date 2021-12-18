package main

import (
	"fmt"
	"github.com/91go/rss2/server/habit"

	code2 "github.com/91go/rss2/server/code"

	"github.com/91go/rss2/utils/resp"

	"github.com/sirupsen/logrus"

	"github.com/91go/rss2/middleware"

	asmr2 "github.com/91go/rss2/server/asmr"
	lf2 "github.com/91go/rss2/server/lf"
	life2 "github.com/91go/rss2/server/life"
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

	// local
	lf := r.Group("/lf")
	lf.GET("/local/:path", lf2.LocalFileRss)
	lf.GET("/local/:path/:sec", lf2.LocalSecDirFileRss)

	// asmr路由
	asmr := r.Group("/asmr")
	asmr.GET("/evc", asmr2.EvcRss)

	code := r.Group("/code")
	code.GET("/gocn/:topic", code2.GoCnRss)

	// life路由分组
	life := r.Group("/life")
	life.GET("/iresearch", life2.IResearchRss)
	life.GET("/weather", life2.WeatherRss)

	// habit
	life.GET("/habit/notify", habit.HabitNotifyRss)
	life.GET("/habit/md", habit.HabitMDRss)
	life.GET("/habit/routine", habit.HabitRoutineRss)

	// porn路由分组
	porn := r.Group("/porn")
	porn.GET("/ysk/:tag", porn2.YskRss)
	porn.GET("/jiuse/:author", porn2.JiuSeRss)
	porn.GET("/dybz/:novel", porn2.DybzRss)
	porn.GET("/dybz/search/:novel", porn2.DybzSearchRss)
	porn.GET("/pornhub/model/:model", porn2.PornhubRss)

	return r
}
