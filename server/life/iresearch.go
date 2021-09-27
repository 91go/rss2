package life

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/91go/rss2/utils/http"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"

	"github.com/sirupsen/logrus"

	"github.com/gogf/gf/os/gtime"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

const (
	BaseURL   = "https://www.iresearch.com.cn/api/products/GetReportList?classId=&fee=0&date=&lastId=&pageSize=9"
	DetailURL = "https://www.iresearch.com.cn/api/Detail/reportM?id=%s&isfree=0"
	LIMIT     = 1
)

// IResearchRss [产业研究报告-艾瑞咨询](https://www.iresearch.com.cn/m/report.shtml)
func IResearchRss(ctx *gin.Context) {
	ret := crawlIResearch()

	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "艾瑞咨询",
		},
		URL: BaseURL,
	}, ret)

	resp.SendXML(ctx, res)
}

func crawlIResearch() []rss.Item {
	body := http.RequestGet(BaseURL)
	res, err := simplejson.NewJson(body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": BaseURL,
			"err": err,
		}).Warn("parse iresearch failed")

		return []rss.Item{}
	}

	iResearch := []rss.Item{}
	rows, err := res.Get("List").Array()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": BaseURL,
			"err": err,
		}).Warn("detail加载失败")

		return []rss.Item{}
	}
	for _, row := range rows[0:LIMIT] {
		if each, ok := row.(map[string]interface{}); ok {
			id := each["NewsId"].(json.Number).String()
			detail := parseDetail(id)

			iResearch = append(iResearch, rss.Item{
				Title:    each["Title"].(string),
				Time:     transTime(each["Uptime"].(string)),
				URL:      each["VisitUrl"].(string),
				Contents: fmt.Sprintf("%s%s", each["Content"].(string), detail),
			})
		}
		continue
	}
	return iResearch
}

// 详情
func parseDetail(id string) (ret string) {
	url := fmt.Sprintf(DetailURL, id)
	body := http.RequestGet(url)
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
