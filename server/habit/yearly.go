package habit

import (
	"fmt"
	"net/http"

	"github.com/91go/rss2/model"

	"github.com/91go/rss2/utils/helper/html"
	"github.com/91go/rss2/utils/helper/time"

	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/os/gtime"
	"github.com/golang-module/carbon"
)

const (
	TwoDaily     = "@2daily"
	ThreeDaily   = "@3daily"
	FourDaily    = "@4daily"
	SixDaily     = "@6daily"
	Weekly       = "@weekly"
	TwoWeekly    = "@2weekly"
	ThreeWeekly  = "@3weekly"
	Monthly      = "@monthly"
	TwoMonthly   = "@2monthly"
	ThreeMonthly = "@3monthly"
	SixMonthly   = "@6monthly"
	Yearly       = "@yearly"
)

func HabitYearlyRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "生活习惯yearly",
		},
		Author:      "lry",
		URL:         GetURL(ctx.Request),
		UpdatedTime: time.GetToday(),
	}, habitFeed())

	resp.SendXML(ctx, res)
}

func habitFeed() []rss.Item {
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
	isSaturday := cb.IsSaturday()
	dayOfYear := cb.DayOfYear()
	dayOfMonth := cb.DayOfMonth()
	weekOfYear := cb.WeekOfYear()
	monthOfYear := cb.MonthOfYear()
	isJanuary := cb.IsJanuary()

	// @2daily
	if cronTime == TwoDaily && ((dayOfYear-1)%2 == 0 || dayOfYear == 1) {
		return true
	}
	// @3daily
	if cronTime == ThreeDaily && ((dayOfYear-1)%3 == 0 || dayOfYear == 1) {
		return true
	}
	if cronTime == FourDaily && ((dayOfYear-1)%4 == 0 || dayOfYear == 1) {
		return true
	}
	// @6daily
	if cronTime == SixDaily && ((dayOfYear-1)%6 == 0 || dayOfYear == 1) {
		return true
	}
	// @weekly
	if cronTime == Weekly && isSaturday {
		return true
	}
	// @2weekly
	if cronTime == TwoWeekly && weekOfYear%2 != 0 && isSaturday {
		return true
	}
	// @3weekly
	if cronTime == ThreeWeekly && weekOfYear%3 != 0 && isSaturday {
		return true
	}
	// @monthly 每月1号
	if cronTime == Monthly && dayOfMonth == 1 {
		return true
	}
	// @2monthly
	if cronTime == TwoMonthly && garray.NewIntArrayFrom([]int{1, 3, 5, 7, 9, 11}).Contains(monthOfYear) && dayOfMonth == 1 {
		return true
	}
	// @3monthly
	if cronTime == ThreeMonthly && garray.NewIntArrayFrom([]int{1, 4, 7, 10}).Contains(monthOfYear) && dayOfMonth == 1 {
		return true
	}
	// @6monthly
	if cronTime == SixMonthly && garray.NewIntArrayFrom([]int{1, 7}).Contains(monthOfYear) && dayOfMonth == 1 {
		return true
	}
	// @yearly
	if cronTime == Yearly && isJanuary && dayOfMonth == 1 {
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
