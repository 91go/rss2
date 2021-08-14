package code

import (
	"fmt"
	"github.com/91go/rss2/utils"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gorilla/feeds"
	"log"
	"time"
)

type HuangZ struct {
	Author string
	Url    string
	Title  string
	Ctx    string
	Time   time.Time
}

var (
	url = "https://blog.huangz.me/"
)

// 用来输出rss
// 全文直出会timeout，所以只打出标题
func HuangZRss(request *ghttp.Request) {

	ass := crawlHuangZ()

	feed := &feeds.Feed{
		Title:       "huangz——黄建宏redis博客",
		Link:        &feeds.Link{Href: url},
		Description: "",
		Author:      &feeds.Author{Name: ass[0].Author},
		Created:     ass[0].Time,
		Updated:     ass[0].Time,
	}

	for _, value := range ass {

		//itemCreateTime, _ := time.Parse("2006-01-02 15:04:05", value["create_time"].(string))
		feed.Add(&feeds.Item{
			Title:       value.Title,
			Link:        &feeds.Link{Href: value.Url},
			Description: "",
			Author:      &feeds.Author{Name: value.Author},
			Content:     value.Ctx,
			Created:     value.Time,
			Updated:     value.Time,
		})
	}

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}

	request.Response.WriteXmlExit(atom)
}

// [huangz/blog — blog.huangz.me](https://blog.huangz.me/#)
func crawlHuangZ() []HuangZ {

	doc := utils.FetchHTML(url)

	wrap := doc.Find(".toctree-l2")

	var param = []HuangZ{}
	wrap.Each(func(i int, selection *query.Selection) {
		articleUrl, _ := selection.Find(".reference").Attr("href")
		title := selection.Find(".reference").Text()

		fullUrl := fmt.Sprintf("%s%s", url, articleUrl)
		//detail := utils.FetchHTML(fullUrl)
		//ctx := detail.Find(".body").Text()

		param = append(param, HuangZ{
			Author: "huangz",
			Url:    fullUrl,
			Title:  title,
			Ctx:    "",
			Time:   time.Now(),
		})
	})

	return param
}
