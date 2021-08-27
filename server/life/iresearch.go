package life

import (
	"encoding/json"
	"fmt"
	"github.com/91go/gofc/fchttp"
	"github.com/91go/rss2/core"
	"github.com/bitly/go-simplejson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"time"
)

var (
	BaseUrl   = "https://www.iresearch.com.cn/api/products/GetReportList?classId=&fee=0&date=&lastId=&pageSize=9"
	DetailUrl = "https://www.iresearch.com.cn/api/Detail/reportM?id=%s&isfree=0"
	LIMIT     = 1
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

	atom := core.Rss(core.Feed{
		Title: "艾瑞咨询——产业研究报告",
		Url:   BaseUrl,
	}, res)

	err := request.Response.WriteXmlExit(atom)
	if err != nil {
		return
	}
}

func crawlIResearch() []core.Feed {

	body := fchttp.RequestGet(BaseUrl)
	res, err := simplejson.NewJson(body)
	if err != nil {
		glog.Errorf("list加载失败 %v", err)
		return []core.Feed{}
	}

	iResearch := []core.Feed{}
	rows, err := res.Get("List").Array()
	for _, row := range rows[0:LIMIT] {
		if each, ok := row.(map[string]interface{}); ok {
			id := each["NewsId"].(json.Number).String()
			detail := parseDetail(id)

			iResearch = append(iResearch, core.Feed{
				Title:    each["Title"].(string),
				Time:     transTime(each["Uptime"].(string)),
				Url:      each["VisitUrl"].(string),
				Contents: fmt.Sprintf("%s%s", each["Content"].(string), detail),
				Pics:     detail,
			})
		}
	}
	return iResearch
}

// 详情
func parseDetail(id string) (ret string) {
	url := fmt.Sprintf(DetailUrl, id)
	body := fchttp.RequestGet(url)
	res, _ := simplejson.NewJson(body)
	total, _ := res.Get("List").GetIndex(0).Get("PagesCount").Int()

	for i := 0; i <= total; i++ {
		pic := fmt.Sprintf("https://pic.iresearch.cn/rimgs/%s/%d.jpg", id, i)

		ret += fmt.Sprintf("<img src=%s>", pic)
	}

	return ret
}

func transTime(str string) time.Time {

	tt, _ := time.Parse("2006/1/02 15:04:05", str)
	return tt
}
