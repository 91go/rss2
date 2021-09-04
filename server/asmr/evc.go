package asmr

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/91go/gofc/fctime"
	"github.com/91go/rss2/core"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/robertkrimen/otto"
)

type Asmr struct {
	Title    string `json:"title"`
	APIURL   string `json:"api_url"`
	OriginID int64  `json:"origin_id"`
	Desc     string `json:"desc"`
	// 真实url
	AudioURL string `json:"audio_url"`
	// 创建时间
	CreateTime time.Time `json:"create_time"`
}

const (
	NzURL   = "https://www.2evc.cn/voiceAppserver//common/sortType?voiceType=1&orderType=0&curPage=1&pageSize=302&cvId=8"
	VoiceJs = `
function unDecrypt(e, n) {
    if ("h" == e.substr(0, 1)) return e;
    function t(e, n, t, o) {
        var r = e,
            a = r.substring(0, n),
            i = r.substring(t);
        return a + o + i
    }
    var o = e.substring(41, 43),
        r = e.substring(46, 48),
        a = parseInt(e.substring(44, 45)),
        i = ["8", "5", "1", "7", "3", "6", "9", "0", "2", "4"],
        s = "";
    i.forEach(function(e, n) {
        a == e && (s = n)
    }),
        e = t(e, 0, 1, "h"),
        e = t(e, 41, 43, r),
        e = t(e, 46, 48, o),
        e = t(e, 44, 45, s);
    var c = "",
        l = "";
    return - 1 == e.indexOf("8.210.46.21") ? n ? (c = "http://149.129.87.151:9090/voice", l = e.substring(32)) : (e && (c = "http://149.129.87.151:9090/test"), l = e.substring(32).replace(/0/g, "1")) : n ? (c = "http://8.210.46.21:9090/voice", l = e.substring(29)) : (e && (c = "http://8.210.46.21:9090/test"), l = e.substring(29).replace(/0/g, "1")),
    c + l
}
`
)

// EvcRss 直接用iina播放url，chrome返回302无法播放
func EvcRss(ctx *gin.Context) {
	ret := parseRequest(NzURL)

	res := core.Rss(&core.Feed{
		Title: "南征付费ASMR音频",
		URL:   NzURL,
	}, ret)

	core.SendXML(ctx, res)
}

//
func parseRequest(url string) []core.Feed {
	body := core.RequestGet(url)
	res, err := simplejson.NewJson(body)
	if err != nil {
		log.Printf("list加载失败 %v", err)
		return []core.Feed{}
	}

	asmr := []core.Feed{}
	rows, err := res.Get("data").Get("pageData").Array()
	if err != nil {
		return []core.Feed{}
	}
	rowws := rows[0:core.LimitItem]
	for _, row := range rowws {
		if each, ok := row.(map[string]interface{}); ok {
			origID, err := each["id"].(json.Number).Int64()
			if err != nil {
				log.Printf("convert origID err %v", err)
			}
			apiURL := fmt.Sprintf("https://www.2evc.cn/voiceAppserver/voice/get?id=%d&telephone=undefined&cvId=8", origID)
			detail := parseDetail(apiURL)

			asmr = append(asmr, core.Feed{
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
	body := core.RequestGet(url)
	res, err := simplejson.NewJson(body)
	if err != nil {
		log.Printf("detail加载失败 %v", err)
		return Asmr{}
	}
	each, err := res.Get("data").Map()
	if err != nil {
		log.Printf("detail加载失败 %v", err)
		return Asmr{}
	}
	fileSrc := each["fileSrc"]
	createTime, err := fctime.MsToTime(each["createDate"].(json.Number).String())
	if err != nil {
		log.Printf("trans time error%v", err.Error())
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
	_, err := vm.Run(VoiceJs)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	const hasOwn = "true"
	call, err := vm.Call("unDecrypt", nil, fileSource, hasOwn)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return call.String()
}
