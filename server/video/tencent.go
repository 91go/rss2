package video

import (
	"fmt"
	"github.com/91go/rss2/utils"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gorilla/feeds"
	"log"
	"time"
)

type TencentVideo struct {
	Url     string
	Episode string
	Time    time.Time
}

var (
	TencentVideoUrl = "https://v.qq.com/detail/m/%s.html"
)

func TencentRss(request *ghttp.Request) {
	tag := request.GetString("tag")
	url := fmt.Sprintf(TencentVideoUrl, tag)

	list, videoName := ParseTencentVideoList(url)

	feed := &feeds.Feed{
		Title: videoName,
		Link:  &feeds.Link{Href: url},
	}
	for _, val := range list {

		feed.Add(&feeds.Item{
			Title: fmt.Sprintf("%s%s", videoName, val.Episode),
			Link:  &feeds.Link{Href: val.Url},
		})
	}

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}

	request.Response.WriteXmlExit(atom)
}

func ParseTencentVideoList(url string) ([]TencentVideo, string) {
	doc := utils.FetchHTML(url)
	videoName := doc.Find(".video_title_cn").Find("a").Text()
	if doc == nil {
		return []TencentVideo{}, videoName
	}
	lists := doc.Find(".mod_episode").Find(".item")

	ret := []TencentVideo{}

	lists.Each(func(i int, selection *query.Selection) {
		href, _ := selection.Find("a").Attr("href")
		episode := selection.Find("a").Find("span").Text()
		isVipEpisode := selection.Find("a").Find("span").HasClass("mark_v")
		if !isVipEpisode {
			ret = append(ret, TencentVideo{
				Url:     href,
				Episode: episode,
			})
		}
	})
	return ret, videoName
}
