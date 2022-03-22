package code

import (
	"fmt"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/helper/time"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/text/gstr"
)

type CodersWayCourse struct {
	CourseName string
	CourseURL  string
	IsUpdating bool          // 是否更新中： true未完结/false完结
	Remark     string        // 难易程度： 初级/中级/高级
	Chapters   []ChapterInfo // 课程章节
}

type ChapterInfo struct {
	URL   string
	Title string
	Intro string
}

type CourseCate struct {
	CateName string
	CateURL  string
}

var (
	BanCourseClass = []string{
		"https://www.jtthink.com/course?c=14", // java/python
		"https://www.jtthink.com/course?c=13", // 前端
		"https://www.jtthink.com/course?c=8",  // PHP
	}
	DefaultClass = "https://www.jtthink.com/course?c=99"
)

const (
	CodersWayBaseURL = "https://www.jtthink.com"
	CodersWayURL     = "https://www.jtthink.com/course"
)

func CodersWayRes(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		URL: CodersWayURL,
		Title: rss.Title{
			Prefix: "程序员在囧途",
			Name:   "xxx",
		},
		UpdatedTime: time.GetToday(),
	}, courseList())

	resp.SendXML(ctx, res)
}

func courseList() []rss.Item {
	doc := gq.FetchHTML(CodersWayURL)
	ret := []rss.Item{}

	cates := []CourseCate{}
	// 获取所有课程分类
	doc.Find(".page-header").Find("a").Each(func(i int, sel *query.Selection) {
		url := CodersWayBaseURL + sel.AttrOr("href", DefaultClass)
		if !garray.NewStrArrayFrom(BanCourseClass).Contains(url) {
			cates = append(cates, CourseCate{
				CateName: sel.Text(),
				CateURL:  url,
			})
		}
	})

	// 分页获取课程分类下的课程以及章节
	for _, cate := range cates {
		gq.FetchHTML(cate.CateURL).Find(".row").Eq(1).Find(".col-md-4").Find(".thumbnail").Find(".caption").Each(func(i int, sel *query.Selection) {
			courseName, _ := sel.Find("h4").Find("a").Attr("title")
			courseURL, _ := sel.Find("h4").Find("a").Attr("href")
			remark := sel.Find("p").Text()
			isUpdating := gstr.Contains(remark, "更新中")

			chapters := parseCourseDetail(courseURL)
			// 剔除所有的已完结课程，只展示更新中的课程
			if isUpdating {
				for _, chapter := range chapters {
					title := fmt.Sprintf("[%s] - %s", courseName, chapter.Title)
					ret = append(ret, rss.Item{
						URL:      chapter.URL,
						Title:    title,
						Contents: chapter.Intro,
						ID:       rss.GenFixedID("coders-way", chapter.URL),
					})
				}
			}
		})
	}

	return ret
}

// 课程详情
func parseCourseDetail(url string) (chapters []ChapterInfo) {
	// https://www.jtthink.com/course/170
	gq.FetchHTML(url).Find(".list-group").Each(func(i int, sel *query.Selection) {
		courseURL, _ := sel.Find(".coursetitle").Find("a").Attr("href")
		title := sel.Find(".coursetitle").Find("a").First().Text()
		intro := sel.Find(".courseintr").Find("p").Text()
		chapters = append(chapters, ChapterInfo{
			URL:   courseURL,
			Title: title,
			Intro: intro,
		})
	})
	return chapters
}
