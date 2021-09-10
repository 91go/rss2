package porn

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/91go/rss2/utils"

	"github.com/gogf/gf/text/gregex"

	"github.com/sirupsen/logrus"

	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/text/gstr"

	"github.com/91go/rss2/core"
	query "github.com/PuerkitoBio/goquery"

	"github.com/gin-gonic/gin"
)

const (
	DybzBaseUrl = "http://m.hongrenxs.net/book/"
)

// DybzRss 第一版主rss源
func DybzRss(ctx *gin.Context) {
	book := ctx.Param("novel")
	url := fmt.Sprintf("%s%s/", DybzBaseUrl, book)

	info, list := dybzList(url)

	res := core.Rss(&info, list)

	core.SendXML(ctx, res)
}

// 某novel的列表
func dybzList(url string) (feed core.Feed, feeds []core.Feed) {
	doc := core.FetchHTML(url)

	wrap := doc.Find(".list_xm").Find("ul").Find("li")
	ret := []core.Feed{}
	wrap.Each(func(i int, selection *query.Selection) {
		title := selection.Find("a").Text()
		novelUrl, _ := selection.Find("a").Attr("href")

		detail, err := novelDetail(novelUrl)
		if err == nil {
			ret = append(ret, core.Feed{
				Title: title,
				URL:   novelUrl,
				Time:  detail,
			})
		} else {
			ret = append(ret, core.Feed{
				Title: title,
				URL:   novelUrl,
				Time:  utils.GetToday(),
			})
		}
	})

	info := dybzInfo(url, doc)

	return info, ret
}

func dybzInfo(url string, doc *query.Document) core.Feed {
	novelName := doc.Find(".cataloginfo").Find("h3").Text()
	author := doc.Find(".infotype").Find("p").Find("a").Text()
	return core.Feed{
		URL:    url,
		Title:  fmt.Sprintf("%s%s", "第一版主-", novelName),
		Author: author,
	}
}

func novelDetail(url string) (time.Time, error) {
	doc := core.FetchHTML(url)
	find := doc.Find(".articlecontent").Find("div").Find("p").Text()

	t := strings.Replace(find, " ", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "&nbsp", "", -1)
	t = strings.Replace(t, " ", "", -1)

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
