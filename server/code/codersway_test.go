package code

import (
	"fmt"
	"testing"

	"github.com/91go/rss2/utils/gq"
	query "github.com/PuerkitoBio/goquery"
)

func TestDetail(t *testing.T) {
	url := "https://www.jtthink.com/course/170"
	doc := gq.FetchHTML(url).Find(".list-group.course")
	length := doc.Length()

	doc.Slice(0, length-1).Each(func(i int, sel *query.Selection) {
		courseURL, _ := sel.Find(".list-group-item.coursetitle.white").Find("a").Attr("href")
		title := sel.Find(".list-group-item.coursetitle.white").Find("a").First().Text()
		intro := sel.Find(".list-group-item.courseintr.white").Find("p").Text()
		// 判断是否为试听课
		isFree := sel.Find(".list-group-item.coursetitle.white").Find("span").Text()
		if isFree == FreeCourse {
			courseURL, _ = sel.Find(".list-group-item.coursetitle.white").Find("a").Eq(1).Attr("href")
			title = sel.Find(".list-group-item.coursetitle.white").Find("a").Eq(1).Text()
		}
		fmt.Println(courseURL, title, intro)
	})
}
