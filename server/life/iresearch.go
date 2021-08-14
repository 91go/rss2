package life

import (
	"encoding/json"
	"fmt"
	"github.com/91go/rss2/utils"
	"github.com/bitly/go-simplejson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gorilla/feeds"
	"log"
	"time"
)

var (
	BaseUrl   = "https://www.iresearch.com.cn/api/products/GetReportList?classId=&fee=0&date=&lastId=&pageSize=9"
	DetailUrl = "https://www.iresearch.com.cn/api/Detail/reportM?id=%s&isfree=0"
	LIMIT     = 6
)

type IResearch struct {
	Title    string
	Time     time.Time
	Url      string
	Describe string
	Pics     string
}

// [产业研究报告-艾瑞咨询](https://www.iresearch.com.cn/m/report.shtml)
func IResearchRss(request *ghttp.Request) {
	res := crawlIResearch()

	feed := &feeds.Feed{
		Title:   "艾瑞咨询——产业研究报告",
		Link:    &feeds.Link{Href: res[0].Url},
		Author:  &feeds.Author{Name: ""},
		Created: res[0].Time,
	}
	for _, value := range res {

		url := value.Url
		feed.Add(&feeds.Item{
			Title:       value.Title,
			Link:        &feeds.Link{Href: url},
			Description: fmt.Sprintf("%s%s", value.Describe, value.Pics),
			Author:      &feeds.Author{Name: ""},
			Created:     value.Time,
		})
	}

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}

	request.Response.WriteXmlExit(atom)
}

func crawlIResearch() []IResearch {

	body := utils.RequestGet(BaseUrl)
	res, err := simplejson.NewJson(body)
	if err != nil {
		glog.Errorf("list加载失败 %v", err)
		return []IResearch{}
	}

	iResearch := []IResearch{}
	rows, err := res.Get("List").Array()
	for _, row := range rows[0:2] {
		if each, ok := row.(map[string]interface{}); ok {
			id := each["NewsId"].(json.Number).String()
			detail := parseDetail(id)
			iResearch = append(iResearch, IResearch{
				Title:    each["Title"].(string),
				Time:     utils.TransTime2(each["Uptime"].(string)),
				Url:      each["VisitUrl"].(string),
				Describe: each["Content"].(string),
				Pics:     detail,
			})
		}
	}
	return iResearch
}

func parseDetail(id string) (ret string) {
	url := fmt.Sprintf(DetailUrl, id)
	body := utils.RequestGet(url)
	res, _ := simplejson.NewJson(body)
	total, _ := res.Get("List").GetIndex(0).Get("PagesCount").Int()

	for i := 0; i <= total; i++ {
		pic := fmt.Sprintf("https://pic.iresearch.cn/rimgs/%s/%d.jpg", id, i)

		ret += fmt.Sprintf("<img src=%s>", pic)
	}

	return ret
}
