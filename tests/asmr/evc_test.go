package asmr

import (
	"github.com/91go/rss2/core"
	"github.com/91go/rss2/server/asmr"
	"github.com/91go/rss2/server/code"
	"github.com/gin-gonic/gin"
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestEvc(t *testing.T) {
	uri := "/code/huangZ"
	router := gin.Default()
	router.GET(uri, code.HuangZRss)

	_, err := core.Get(uri, router)
	assert.Nil(t, err)
	//assert.Len(t, )
	//fmt.Println(string(res))
}

func TestVoice(t *testing.T) {

	vm := otto.New()
	res, err := vm.Run(asmr.VoiceJs())
	if err != nil {
		log.Println(err.Error())
	}
	t.Log(res)

	fileSource := "5ttp://8.210.46.21:9090/voice/60000000765.m97973p3"
	hasOwn := "true"
	call, err := vm.Call("unDecrypt", nil, fileSource, hasOwn)
	if err != nil {
		return
	}
	t.Log(call)
}

//
//func TestApiUrl(t *testing.T) {
//	all, err := dao.Asmr.All()
//	if err != nil {
//		return
//	}
//	for _, one := range all {
//		//originId, _ := one.OriginId
//		apiUrl := fmt.Sprintf("https://www.2evc.cn/voiceAppserver/voice/get?id=%d&telephone=undefined&cvId=8", one.OriginId)
//		_, err := dao.Asmr.Data("api_url", apiUrl).Where("origin_id", one.OriginId).Update()
//		if err != nil {
//			t.Errorf("update failed: %d", one.OriginId)
//			//return
//		}
//	}
//}
//
//func TestParseDetail(t *testing.T) {
//	ParseDetail()
//}
//
//func TestSimpleJson(t *testing.T) {
//	apiUrl := fmt.Sprintf("https://www.2evc.cn/voiceAppserver/voice/get?id=%d&telephone=undefined&cvId=8", 766)
//	body := RequestGet(apiUrl)
//	res, err := simplejson.NewJson(body)
//	if err != nil {
//		fmt.Printf("%v\n", err)
//		return
//	}
//	rows, err := res.Get("data").Map()
//	i := rows["id"]
//	t.Log(i)
//}
//
//func TestDownloadAudio(t *testing.T) {
//	DownloadAudio()
//}
//
//// fix之前is_download错误的问题
//func TestUpdateIsDownloadFlag(t *testing.T) {
//
//	dirPath := "/Users/luruiyang/Downloads/nz"
//	dir, err := ioutil.ReadDir(dirPath)
//	if err != nil {
//		return
//	}
//	filenames := []string{}
//	for _, file := range dir {
//
//		filename := file.Name()
//		filenames = append(filenames, filename)
//	}
//
//	all, err := dao.Asmr.Fields("code", "title").Where("is_download", 0).All()
//	if err != nil {
//		return
//	}
//	for _, one := range all {
//		dao.Asmr.Where("code", one.Code).Data("is_download", 1).Update()
//	}
//}
