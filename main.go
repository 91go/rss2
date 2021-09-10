package main

import (
	"fmt"

	"github.com/91go/rss2/middleware"

	asmr2 "github.com/91go/rss2/server/asmr"
	code2 "github.com/91go/rss2/server/code"
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

	// asmr路由
	asmr := r.Group("/asmr")
	asmr.GET("evc", asmr2.EvcRss)

	// code路由分组
	code := r.Group("/code")
	code.GET("/huangZ", code2.HuangZRss)

	// life路由分组
	life := r.Group("/life")
	life.GET("/iresearch", life2.IResearchRss)
	life.GET("/weather", life2.WeatherRss)

	// porn路由分组
	porn := r.Group("/porn")
	porn.GET("/ysk/:tag", porn2.YskRss)
	porn.GET("/jiuse/:author", porn2.JiuSeRss)
	porn.GET("/dybz/:novel", porn2.DybzRss)

	return r
}
