package asmr

// todo 添加所有的httptest测试

// func TestApiURL(t *testing.T) {
// 	all, err := dao.Asmr.All()
// 	if err != nil {
// 		return
// 	}
// 	for _, one := range all {
// 		//originId, _ := one.OriginId
// 		apiURL := fmt.Sprintf("https://www.2evc.cn/voiceAppserver/voice/get?id=%d&telephone=undefined&cvId=8", one.OriginId)
// 		_, err := dao.Asmr.Data("api_url", apiURL).Where("origin_id", one.OriginId).Update()
// 		if err != nil {
// 			t.Errorf("update failed: %d", one.OriginId)
// 			//return
// 		}
// 	}
// }

// func TestParseDetail(t *testing.T) {
// 	ParseDetail()
// }
//
// func TestSimpleJson(t *testing.T) {
// 	apiURL := fmt.Sprintf("https://www.2evc.cn/voiceAppserver/voice/get?id=%d&telephone=undefined&cvId=8", 766)
// 	body := RequestGet(apiURL)
// 	res, err := simplejson.NewJson(body)
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 		return
// 	}
// 	rows, err := res.Get("data").Map()
// 	i := rows["id"]
// 	t.Log(i)
// }
//
// func TestDownloadAudio(t *testing.T) {
// 	DownloadAudio()
// }
//
// // fix之前is_download错误的问题
// func TestUpdateIsDownloadFlag(t *testing.T) {
//
// 	dirPath := "/Users/luruiyang/Downloads/nz"
// 	dir, err := ioutil.ReadDir(dirPath)
// 	if err != nil {
// 		return
// 	}
// 	filenames := []string{}
// 	for _, file := range dir {
//
// 		filename := file.Name()
// 		filenames = append(filenames, filename)
// 	}
//
// 	all, err := dao.Asmr.Fields("code", "title").Where("is_download", 0).All()
// 	if err != nil {
// 		return
// 	}
// 	for _, one := range all {
// 		dao.Asmr.Where("code", one.Code).Data("is_download", 1).Update()
// 	}
// }
