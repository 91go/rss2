package main

import (
	asmr2 "github.com/91go/rss2/server/asmr"
	code2 "github.com/91go/rss2/server/code"
	life2 "github.com/91go/rss2/server/life"
	porn2 "github.com/91go/rss2/server/porn"
	"github.com/gin-gonic/gin"
)

//func main() {
//	s := g.Server()
//	s.SetPort(8090)
//	s.Run()
//}
//
//func init() {
//	s := g.Server()
//
//	s.Group("/asmr", func(group *ghttp.RouterGroup) {
//		group.GET("/evc", asmr.EvcRss)
//	})
//
//	s.Group("/code", func(group *ghttp.RouterGroup) {
//		group.GET("/huangZ", code.HuangZRss)
//		// 巢大博客
//		group.GET("/xargin", code.XarginRss)
//	})
//
//	s.Group("/life", func(group *ghttp.RouterGroup) {
//		// 艾瑞咨询
//		group.GET("/iresearch", life.IResearchRss)
//		group.GET("/weather", life.WeatherRss)
//	})
//	s.Group("/porn", func(group *ghttp.RouterGroup) {
//		//group.GET("/dybz", porn.DybzRss)
//		group.GET("/ysk/{tag}", porn.YskRss)
//		group.GET("/dybz/{novel}", porn.DybzRss)
//	})
//}

func main() {
	r := gin.Default()

	err := r.Run(":8090")
	if err != nil {
		return
	}
}

func init() {
	r := gin.Default()

	asmr := r.Group("/asmr")
	{
		asmr.GET("evc", asmr2.EvcRss)
	}
	code := r.Group("/code")
	{
		code.GET("/huangZ", code2.HuangZRss)
		//code.GET("/xargin", code2.XarginRss)
	}
	life := r.Group("/life")
	{
		life.GET("/iresearch", life2.IResearchRss)
		life.GET("/weather", life2.WeatherRss)
	}
	porn := r.Group("/porn")
	{
		porn.GET("/ysk/{tag}", porn2.YskRss)
		//porn.GET("/dybz/{novel}", porn2.DybzRss)
	}
}
