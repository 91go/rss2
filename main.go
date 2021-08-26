package main

import (
	"github.com/91go/rss2/server/asmr"
	"github.com/91go/rss2/server/code"
	"github.com/91go/rss2/server/life"
	"github.com/91go/rss2/server/porn"
	"github.com/91go/rss2/server/video"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.SetPort(8090)
	s.Run()
}

func init() {
	s := g.Server()

	s.Group("/asmr", func(group *ghttp.RouterGroup) {
		group.GET("/evc", asmr.EvcRss)
	})

	s.Group("/code", func(group *ghttp.RouterGroup) {
		group.GET("/huangZ", code.HuangZRss)
		// 巢大博客
		group.GET("/xargin", code.XarginRss)
	})

	s.Group("/life", func(group *ghttp.RouterGroup) {
		// 艾瑞咨询
		group.GET("/iresearch", life.IResearchRss)
		group.GET("/weather", life.WeatherRss)
	})
	s.Group("/porn", func(group *ghttp.RouterGroup) {
		//group.GET("/dybz", porn.DybzRss)
		group.GET("/ysk/{tag}", porn.YskRss)
		group.GET("/dybz/{novel}", porn.DybzRss)
	})

	s.Group("/video", func(group *ghttp.RouterGroup) {
		// 非vip腾讯视频
		group.GET("/tencent/{tag}", video.TencentRss)
	})
}
