package life

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/91go/rss2/utils"

	"github.com/sirupsen/logrus"

	"github.com/gogf/gf/os/gtime"

	"github.com/91go/rss2/core"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

const (
	BaseURL   = "https://www.iresearch.com.cn/api/products/GetReportList?classId=&fee=0&date=&lastId=&pageSize=9"
	DetailURL = "https://www.iresearch.com.cn/api/Detail/reportM?id=%s&isfree=0"
	LIMIT     = 1
)

type IResearch struct {
	Title    string
	Time     time.Time
	URL      string
	Describe string
	Pics     string
}

// IResearchRss [产业研究报告-艾瑞咨询](https://www.iresearch.com.cn/m/report.shtml)
func IResearchRss(ctx *gin.Context) {
	ret := crawlIResearch()

	res := core.Rss(&core.Feed{
		Title: "艾瑞咨询——产业研究报告",
		URL:   BaseURL,
	}, ret)

	core.SendXML(ctx, res)
}

func crawlIResearch() []core.Feed {
	body := utils.RequestGet(BaseURL)
	res, err := simplejson.NewJson(body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": BaseURL,
			"err": err,
		}).Warn("parse iresearch failed")

		return []core.Feed{}
	}

	iResearch := []core.Feed{}
	rows, err := res.Get("List").Array()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": BaseURL,
			"err": err,
		}).Warn("detail加载失败")

		return []core.Feed{}
	}
	for _, row := range rows[0:LIMIT] {
		if each, ok := row.(map[string]interface{}); ok {
			id := each["NewsId"].(json.Number).String()
			detail := parseDetail(id)

			iResearch = append(iResearch, core.Feed{
				Title:    each["Title"].(string),
				Time:     transTime(each["Uptime"].(string)),
				URL:      each["VisitUrl"].(string),
				Contents: fmt.Sprintf("%s%s", each["Content"].(string), detail),
				Pics:     detail,
			})
		}
		continue
	}
	return iResearch
}

// 详情
func parseDetail(id string) (ret string) {
	url := fmt.Sprintf(DetailURL, id)
	body := utils.RequestGet(url)
	res, _ := simplejson.NewJson(body)
	total, _ := res.Get("List").GetIndex(0).Get("PagesCount").Int()

	for i := 0; i <= total; i++ {
		pic := fmt.Sprintf("https://pic.iresearch.cn/rimgs/%s/%d.jpg", id, i)

		ret += fmt.Sprintf("<img src=%s>", pic)
	}

	return ret
}

func transTime(str string) time.Time {
	format, err := gtime.StrToTimeFormat(str, "Y/n/d H:i:s")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"time": str,
			"err":  err,
		}).Warn("transTime failed")
		return time.Time{}
	}
	return format.Time
}
