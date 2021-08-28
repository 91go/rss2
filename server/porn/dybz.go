package porn

//func DybzRss(ctx *gin.Context) {
//
//	novel := ctx.GetString("novel")
//
//	list, err := GetRssByTag(novel)
//	if err != nil {
//		panic(err)
//	}
//
//	ass := list.List()
//	//if len(ass) == 0 {
//	//	CrawlLatestUrl(novelUrl)
//	//	//time.Sleep(time.Second * 20)
//	//}
//	feedCreateTime, _ := time.Parse("2006-01-02 15:04:05", ass[0]["create_time"].(string))
//
//	feed := &feeds.Feed{
//		Title:       ass[0]["novel_name"].(string),
//		Link:        &feeds.Link{Href: ass[0]["novel_url"].(string)},
//		Description: "第一版主",
//		Author:      &feeds.Author{Name: "", Email: ""},
//		Created:     feedCreateTime,
//		Updated:     feedCreateTime,
//	}
//
//	for _, value := range ass {
//
//		itemCreateTime, _ := time.Parse("2006-01-02 15:04:05", value["create_time"].(string))
//		feed.Add(&feeds.Item{
//
//			Title:       value["chapter_name"].(string),
//			Link:        &feeds.Link{Href: value["chapter_url"].(string)},
//			Description: "",
//			Created:     itemCreateTime,
//			Updated:     itemCreateTime,
//		})
//	}
//
//	res, err := feed.ToAtom()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	ctx.Data(200, "application/xml; charset=utf-8", []byte(res))
//}
//
//func GetRssByTag(flag string) (gdb.Result, error) {
//	list, err := g.DB().GetAll("select * from dybz where novel_flag = ? order by --chapter_flag desc", flag)
//	return list, err
//}
