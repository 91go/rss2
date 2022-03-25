package habit

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/91go/rss2/utils/helper/iter"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/text/gstr"

	"github.com/91go/rss2/model"

	"github.com/91go/rss2/utils/helper/html"
	"github.com/91go/rss2/utils/helper/time"

	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
	"github.com/golang-module/carbon"
)

var numberReg = regexp.MustCompile(`\d+`)

func HabitYearlyRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "生活习惯yearly",
		},
		Author:      "lry",
		URL:         GetURL(ctx.Request),
		UpdatedTime: time.GetToday(),
	}, yearly())

	resp.SendXML(ctx, res)
}

func yearly() []rss.Item {
	ret := []rss.Item{}
	m := model.Yearly{}
	items, _ := m.FindAll()

	for _, item := range items {
		if CheckCron(item.Cron, carbon.Now()) {
			title := fmt.Sprintf("[%s] - [%s] - %s", item.Prefix, gtime.Date(), item.Task)
			ret = append(ret, rss.Item{
				Title:       title,
				Contents:    html.Md2HTML(item.Remark),
				UpdatedTime: time.GetToday(),
				ID:          rss.GenDateID("habit-notify", item.Task),
			})
		}
	}

	return ret
}

func CheckCron(cronTime string, cb carbon.Carbon) bool {
	dayOfYear := cb.DayOfYear()
	dayOfMonth := cb.DayOfMonth()
	weekOfYear := cb.WeekOfYear()
	monthOfYear := cb.MonthOfYear()

	isSaturday := cb.IsSaturday()
	isJanuary := cb.IsJanuary()
	isMatched, number := ExtractTimeNumber(cronTime)

	// 判断daily
	if gstr.Contains(cronTime, "daily") && ((dayOfYear-1)%number == 0 || dayOfYear == 1) {
		return true
	}
	// 判断weekly
	if gstr.Contains(cronTime, "weekly") && isSaturday && (weekOfYear%number != 0 || !isMatched) {
		return true
	}
	// 判断monthly
	months := GetMonths(number)
	isContains := garray.NewIntArrayFrom(months).Contains(monthOfYear)
	if gstr.Contains(cronTime, "monthly") && dayOfMonth == 1 && isContains {
		return true
	}
	// @yearly
	if gstr.Contains(cronTime, "yearly") && isJanuary && dayOfMonth == 1 {
		return true
	}

	return false
}

func GetURL(r *http.Request) string {
	scheme := "http://"

	if r.TLS != nil {
		scheme = "https://"
	}
	return scheme + r.Host + r.RequestURI
}

// ExtractTimeNumber 正则提取数字
func ExtractTimeNumber(t string) (isMatched bool, number int) {
	isMatched = numberReg.MatchString(t)
	if !isMatched {
		return isMatched, 1
	}
	number, _ = strconv.Atoi(numberReg.FindString(t))
	return isMatched, number
}

// GetMonths 间隔取值
func GetMonths(step int) (res []int) {
	for v := range iter.N(1, carbon.MonthsPerYear+1, step) {
		res = append(res, v)
	}
	return
}
