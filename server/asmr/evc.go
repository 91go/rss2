package asmr

import (
	"encoding/json"
	"fmt"
	"github.com/91go/gofc/fchttp"
	"github.com/91go/gofc/fctime"
	"github.com/91go/rss2/core"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/robertkrimen/otto"
	"log"
	"time"
)

type Asmr struct {
	Title    string `json:"title"`
	ApiUrl   string `json:"api_url"`
	OriginId int64  `json:"origin_id"`
	Desc     string `json:"desc"`
	// 真实url
	AudioUrl string `json:"audio_url"`
	// 创建时间
	CreateTime time.Time `json:"create_time"`
}

var (
	NzURL = "https://www.2evc.cn/voiceAppserver//common/sortType?voiceType=1&orderType=0&curPage=1&pageSize=302&cvId=8"
	// 机器性能不行，不要开太大，否则502
	LIMIT = 3
)

// 直接用iina播放url，chrome返回302无法播放
func EvcRss(ctx *gin.Context) {

	ret := parseRequest(NzURL)

	res := core.Rss(core.Feed{
		Title: "南征付费ASMR音频",
		Url:   NzURL,
	}, ret)

	ctx.Data(200, "application/xml; charset=utf-8", []byte(res))
}

//
func parseRequest(url string) []core.Feed {
	body := fchttp.RequestGet(url)
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
	rowws := rows[0:LIMIT]
	for _, row := range rowws {

		if each, ok := row.(map[string]interface{}); ok {

			origId, err := each["id"].(json.Number).Int64()
			if err != nil {
				log.Printf("convert origId err %v", err)
			}
			apiUrl := fmt.Sprintf("https://www.2evc.cn/voiceAppserver/voice/get?id=%d&telephone=undefined&cvId=8", origId)
			detail := parseDetail(apiUrl)

			asmr = append(asmr, core.Feed{
				Title:    each["name"].(string),
				Url:      detail.AudioUrl,
				Contents: fmt.Sprintf("%s\n%s", detail.AudioUrl, detail.Desc),
				Time:     detail.CreateTime,
			})
		}
	}

	return asmr
}

// 解析详情页
func parseDetail(url string) Asmr {
	body := fchttp.RequestGet(url)
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
		AudioUrl:   originAudioUrl(fileSrc.(string)),
		CreateTime: createTime,
	}
}

// 获取音频的真实url
func originAudioUrl(fileSource string) string {

	vm := otto.New()
	_, err := vm.Run(VoiceJs())
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	hasOwn := "true"
	call, err := vm.Call("unDecrypt", nil, fileSource, hasOwn)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return call.String()
}

func VoiceJs() string {
	return `
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
}
