package porn

import (
	"github.com/91go/rss2/core"
	query "github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	novels = []string{}
	latest = []string{
		// 心路难平2
		//"http://www.dybz88.com/22/22417/",
		// 母上攻略
		"https://www.dybz7.net/14/14247/",
		// 新婚
		"http://www.dybz88.com/11/11812/",
		// 六朝燕歌行
		"http://www.dybz88.com/6/6532/",
		// 枕上余温
		"http://www.dybz77.com/12/12057/",
	}
	search = []string{
		// 逍遥小散仙
		"http://www.dybz88.com/21/21845/",
	}
	base = "https://www.dybz7.net"
)

type Page struct {
	//Page string
	Url string
}

type Dybz struct {
	NovelName   string
	Author      string
	NovelFlag   string
	NovelUrl    string
	ChapterName string
	ChapterUrl  string
	ChapterFlag string
	CreateTime  string
	UpdateTime  string
}

//func CrawlLatest() {
//
//	if len(latest) == 0 {
//		return
//	}
//	for _, url := range latest {
//		glog.Infof("正在爬取: %s", url)
//		novel := LatestNovel(url)
//
//		save, err := dao.Dybz.Data(novel).Insert()
//		if err != nil {
//			glog.Error("insert failed: ", err.Error())
//			//return
//		}
//
//		glog.Error("insert success, result: ", save)
//	}
//}

//func CrawlLatestUrl(url string) {
//	novel := LatestNovel(url)
//
//	save, err := dao.Dybz.Data(novel).Insert()
//	if err != nil {
//		glog.Error("insert failed: ", err.Error())
//		//return
//	}
//	glog.Error("insert success, result: ", save)
//}

func GetNovel(url string) []Dybz {
	doc := core.FetchHTMLNotUtf8(url, "gbk")

	novelName := doc.Find("h1").Text()
	author := doc.Find(".info").Text()
	novelFlag := strings.Split(url, "/")
	flag := novelFlag[len(novelFlag)-2]

	// 获取该小说总页数，循环
	pages := ParsePages(doc, url)
	ret := []Dybz{}

	for _, page := range pages {

		detail := core.FetchHTMLNotUtf8(page.Url, "gbk")
		wrap := detail.Find(".list li")

		wrap.Each(func(i int, selection *query.Selection) {
			chapterName := selection.Text()
			chapterUrl, _ := selection.Find("a").Attr("href")
			chapterFlag := strings.Split(chapterUrl, "/")
			flag2 := strings.Trim(chapterFlag[len(chapterFlag)-1], ".html")

			ret = append(ret, Dybz{
				NovelName:   novelName,
				Author:      author,
				NovelFlag:   flag,
				NovelUrl:    url,
				ChapterName: chapterName,
				ChapterUrl:  base + chapterUrl,
				ChapterFlag: flag2,
				CreateTime:  time.Now().String(),
				UpdateTime:  time.Now().String(),
			})
		})
	}
	return ret
}

func ParsePages(doc *query.Document, url string) (pages []Page) {

	attr, exists := doc.Find(".pagelistbox .page .endPage").Attr("href")
	if !exists {
		log.Fatalln("not exist end page")
	}
	s := strings.Trim(attr, "/")
	e := strings.Split(s, "_")
	ee := e[len(e)-1]
	end, _ := strconv.Atoi(ee)

	for i := 1; i <= end; i++ {
		nth := replaceNth(url, "/", "_"+strconv.Itoa(i)+"/", 5)
		pages = append(pages, Page{Url: nth})
	}

	return pages
}

func LatestNovel(url string) []Dybz {
	doc := core.FetchHTMLNotUtf8(url, "gbk")

	novelName := doc.Find("h1").Text()
	author := doc.Find(".info").Text()
	novelFlag := strings.Split(url, "/")
	flag := novelFlag[len(novelFlag)-2]

	ret := []Dybz{}
	wrap := doc.Find(".chapter-list .bd .list").First()

	wrap.Find("li").Each(func(i int, selection *query.Selection) {
		chapterName := selection.Text()
		chapterUrl, _ := selection.Find("a").Attr("href")
		chapterFlag := strings.Split(chapterUrl, "/")
		flag2 := strings.Trim(chapterFlag[len(chapterFlag)-1], ".html")

		ret = append(ret, Dybz{
			NovelName:   novelName,
			Author:      author,
			NovelFlag:   flag,
			NovelUrl:    url,
			ChapterName: chapterName,
			ChapterUrl:  base + chapterUrl,
			ChapterFlag: flag2,
			CreateTime:  time.Now().String(),
			UpdateTime:  time.Now().String(),
		})
	})

	return ret
}

// Replace the nth occurrence of old in s by new.
func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}
