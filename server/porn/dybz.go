package porn

import (
	"errors"
	"fmt"
	"time"

	"github.com/91go/rss2/core/resp"
	"github.com/91go/rss2/core/rss"

	"github.com/91go/rss2/core/gq"

	"github.com/gogf/gf/frame/g"

	"github.com/91go/rss2/utils"

	"github.com/gogf/gf/text/gregex"

	"github.com/sirupsen/logrus"

	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/text/gstr"

	query "github.com/PuerkitoBio/goquery"

	"github.com/gin-gonic/gin"
)

const (
	DybzBaseUrl   = "http://m.hongrenxs.net"
	DybzBookUrl   = "http://m.hongrenxs.net/book/"
	DybzSearchUrl = "http://m.hongrenxs.net/s.html"
)

// DybzRss 第一版主rss源
func DybzRss(ctx *gin.Context) {
	novel := ctx.Param("novel")
	url := fmt.Sprintf("%s%s/", DybzBookUrl, novel)

	info, list := dybzList(url)
	res := rss.Rss(&info, list)

	resp.SendXML(ctx, res)
}

// DybzSearchRss 搜索某小说
func DybzSearchRss(ctx *gin.Context) {
	novel := ctx.Param("novel")
	m := g.Map{
		"s":    novel,
		"type": "articlename",
	}

	doc := gq.PostHTML(DybzSearchUrl, m)
	url, exists := doc.Find(".searchresult").Find(".sone").Find("a").Attr("href")
	if !exists {
		logrus.WithFields(utils.Fields(url, errors.New("not exist novel")))
		return
	}

	// 根据id获取最新小说，返回小说url
	info, list := dybzList(fmt.Sprintf("%s%s", DybzBaseUrl, url))
	res := rss.Rss(&info, list)

	resp.SendXML(ctx, res)
}

// 某novel的列表
func dybzList(url string) (feed rss.Feed, feeds []rss.Feed) {
	doc := gq.FetchHTML(url)

	wrap := doc.Find(".list_xm").Find("ul").Find("li")
	ret := []rss.Feed{}
	wrap.Each(func(i int, selection *query.Selection) {
		title := selection.Find("a").Text()
		novelUrl, _ := selection.Find("a").Attr("href")

		detail, err := novelDetail(novelUrl)
		if err == nil {
			ret = append(ret, rss.Feed{
				Title: title,
				URL:   novelUrl,
				Time:  detail,
			})
		} else {
			ret = append(ret, rss.Feed{
				Title: title,
				URL:   novelUrl,
				Time:  utils.GetToday(),
			})
		}
	})

	info := dybzInfo(url, doc)

	return info, ret
}

func dybzInfo(url string, doc *query.Document) rss.Feed {
	novelName := doc.Find(".cataloginfo").Find("h3").Text()
	author := doc.Find(".infotype").Find("p").Find(gq.LabelA).Text()
	return rss.Feed{
		URL:    url,
		Title:  fmt.Sprintf("%s%s", "第一版主-", novelName),
		Author: author,
	}
}

func novelDetail(url string) (time.Time, error) {
	doc := gq.FetchHTML(url)
	find := doc.Find(".articlecontent").Find("div").Find("p").Text()

	t := utils.TrimBlank(find)
	if !gstr.Contains(t, "年") {
		return time.Time{}, errors.New("not contains time")
	}

	st, err := gregex.MatchString("(?U)([0-9]+)年([0-9]+)月([0-9]+)日", t)
	if err != nil {
		return time.Time{}, err
	}
	ss := gstr.Join(st[1:], "-")
	tt, err := gtime.StrToTimeFormat(ss, "Y-n-j")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"time": t,
			"err":  err,
		}).Warn("trans time failed")
		return time.Time{}, err
	}

	return tt.Time, nil
}
