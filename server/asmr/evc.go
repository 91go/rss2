package asmr

import (
	"encoding/json"
	"fmt"
	"github.com/91go/gofc/fchttp"
	"github.com/91go/gofc/fctime"
	"github.com/bitly/go-simplejson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gorilla/feeds"
	"github.com/robertkrimen/otto"
	"io/ioutil"
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
	URL = "https://www.2evc.cn/voiceAppserver//common/sortType?voiceType=1&orderType=0&curPage=1&pageSize=302&cvId=8"
	// 机器性能不行，不要开太大，否则502
	LIMIT = 2
)

// 直接用iina播放url，chrome返回302无法播放
func EvcRss(request *ghttp.Request) {

	ret := parseRequest()
	feed := &feeds.Feed{
		Title:   "南征付费ASMR音频",
		Link:    &feeds.Link{Href: URL},
		Author:  &feeds.Author{Name: ""},
		Created: ret[0].CreateTime,
	}

	for _, value := range ret {
		feed.Add(&feeds.Item{
			Title:       value.Title,
			Link:        &feeds.Link{Href: value.AudioUrl},
			Description: value.Desc,
			Created:     value.CreateTime,
		})
	}
	atom, err := feed.ToAtom()
	if err != nil {
		glog.Error(err)
		err := request.Response.WriteXmlExit(atom)
		if err != nil {
			return
		}
	}

	err = request.Response.WriteXmlExit(atom)
	if err != nil {
		return
	}
}

//
func parseRequest() []Asmr {
	body := fchttp.RequestGet(URL)
	res, err := simplejson.NewJson(body)
	if err != nil {
		glog.Errorf("list加载失败 %v", err)
		return []Asmr{}
	}

	asmr := []Asmr{}
	rows, err := res.Get("data").Get("pageData").Array()
	rowws := rows[0:LIMIT]
	for _, row := range rowws {

		if each, ok := row.(map[string]interface{}); ok {

			origId, err := each["id"].(json.Number).Int64()
			if err != nil {
				glog.Error("convert origId err %v", err)
			}
			apiUrl := fmt.Sprintf("https://www.2evc.cn/voiceAppserver/voice/get?id=%d&telephone=undefined&cvId=8", origId)
			detail := parseDetail(apiUrl)

			asmr = append(asmr, Asmr{
				Title:      each["name"].(string),
				OriginId:   origId,
				ApiUrl:     apiUrl,
				Desc:       detail.Desc,
				AudioUrl:   detail.AudioUrl,
				CreateTime: detail.CreateTime,
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
		glog.Errorf("detail加载失败 %v", err)
		return Asmr{}
	}
	each, err := res.Get("data").Map()
	fileSrc := each["fileSrc"]
	createTime, err := fctime.MsToTime(each["createDate"].(json.Number).String())
	if err != nil {
		glog.Errorf("trans time error%v", err.Error())
	}

	return Asmr{
		Desc:       each["voiceDesc"].(string),
		AudioUrl:   originAudioUrl(fileSrc.(string)),
		CreateTime: createTime,
	}
}

// 获取音频的真实url
func originAudioUrl(fileSource string) string {

	jsFile := "./public/js/voice.js"
	bytes, err := ioutil.ReadFile(jsFile)
	if err != nil {
		glog.Error(err.Error())
		return ""
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err != nil {
		glog.Error(err.Error())
		return ""
	}

	hasOwn := "true"
	call, err := vm.Call("unDecrypt", nil, fileSource, hasOwn)
	if err != nil {
		glog.Error(err.Error())
		return ""
	}
	return call.String()
}

// 下载音频
//func DownloadAudio() {
//	all, err := dao.Asmr.Fields("code", "title", "audio_url").Where("is_download", 0).All()
//	if err != nil {
//		return
//	}
//
//	for _, url := range all {
//		audioUrl := url.AudioUrl
//		glog.Info(audioUrl)
//		body := RequestGet(audioUrl)
//
//		err := ioutil.WriteFile(fmt.Sprintf("%s/%s.mp3", "/Users/luruiyang/Downloads/nz", url.Title), body, 0666)
//		if err != nil {
//			glog.Errorf("download failed: %s", audioUrl)
//			return
//		}
//
//		_, err = dao.Asmr.Where("code", url.Code).Data("is_download", 1).Update()
//		if err != nil {
//			glog.Errorf("update download flag error: %s", audioUrl)
//		}
//		glog.Infof("download success: %s", audioUrl)
//	}
//}
