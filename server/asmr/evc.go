package asmr

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"github.com/91go/rss2/core/resp"
	"github.com/91go/rss2/core/rss"

	"github.com/91go/rss2/utils"

	"github.com/sirupsen/logrus"

	"github.com/gogf/gf/os/gfile"

	"github.com/91go/gofc/fctime"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/robertkrimen/otto"
)

type Asmr struct {
	Title      string    `json:"title"`
	APIURL     string    `json:"api_url"`
	OriginID   int64     `json:"origin_id"`
	Desc       string    `json:"desc"`
	AudioURL   string    `json:"audio_url"`   // 真实url
	CreateTime time.Time `json:"create_time"` // 创建时间
}

const (
	NzURL = "https://www.2evc.cn/voiceAppserver//common/sortType?voiceType=1&orderType=0&curPage=1&pageSize=302&cvId=8"
)

// EvcRss 直接用iina播放url，chrome返回302无法播放
func EvcRss(ctx *gin.Context) {
	ret := parseRequest(NzURL)

	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "asmr-南征",
		},
		URL: NzURL,
	}, ret)

	resp.SendXML(ctx, res)
}

func parseRequest(url string) []rss.Item {
	body := utils.RequestGet(url)
	res, err := simplejson.NewJson(body)
	if err != nil {
		logrus.WithFields(utils.Fields(url, err)).Error("list load failed")
		return []rss.Item{}
	}

	asmr := []rss.Item{}
	rows, err := res.Get("data").Get("pageData").Array()
	if err != nil {
		return []rss.Item{}
	}
	rowws := rows[0:rss.LimitItem]
	for _, row := range rowws {
		if each, ok := row.(map[string]interface{}); ok {
			origID, err := each["id"].(json.Number).Int64()
			if err != nil {
				logrus.WithFields(utils.Fields(url, err)).Error("convert origID err")
			}
			apiURL := fmt.Sprintf("https://www.2evc.cn/voiceAppserver/voice/get?id=%d&telephone=undefined&cvId=8", origID)
			detail := parseDetail(apiURL)

			asmr = append(asmr, rss.Item{
				Title:    each["name"].(string),
				URL:      detail.AudioURL,
				Contents: fmt.Sprintf("%s\n%s", detail.AudioURL, detail.Desc),
				Time:     detail.CreateTime,
			})
		}

		continue
	}

	return asmr
}

// 解析详情页
func parseDetail(url string) Asmr {
	body := utils.RequestGet(url)
	res, err := simplejson.NewJson(body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": url,
			"err": err,
		}).Warn("详情页加载失败")
		return Asmr{}
	}
	each, err := res.Get("data").Map()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": url,
			"err": err,
		}).Warn("详情页加载失败")
		return Asmr{}
	}
	fileSrc := each["fileSrc"]
	createTime, err := fctime.MsToTime(each["createDate"].(json.Number).String())
	if err != nil {
		logrus.WithFields(utils.Fields(url, err)).Warn("trans time error")
		return Asmr{}
	}

	return Asmr{
		Desc:       each["voiceDesc"].(string),
		AudioURL:   originAudioURL(fileSrc.(string)),
		CreateTime: createTime,
	}
}

// 获取音频的真实url
func originAudioURL(fileSource string) string {
	vm := otto.New()
	_, err := vm.Run(getPublicFile())
	if err != nil {
		logrus.WithFields(utils.Fields(fileSource, err)).Warn("otto parse js file failed")
		return ""
	}

	const hasOwn = "true"
	call, err := vm.Call("unDecrypt", nil, fileSource, hasOwn)
	if err != nil {
		logrus.WithFields(utils.Fields(fileSource, err)).Warn("otto decrypt failed")
		return ""
	}
	return call.String()
}

// voice.js
func getPublicFile() string {
	abs, err := filepath.Abs("./public/js/voice.js")
	if err != nil {
		logrus.WithFields(utils.Fields("", err)).Warn("voice.js not found")
		return ""
	}
	contents := gfile.GetContents(abs)
	return contents
}
