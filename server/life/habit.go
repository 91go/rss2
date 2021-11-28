package life

import (
	"fmt"

	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/container/garray"
	"github.com/golang-module/carbon"
)

type Notification struct {
	Prefix string // 前缀
	Task   string // 任务
	Cron   string // 执行时间
}

// 用rss代替"提醒事项APP"的原因是，
func HabitRss(ctx *gin.Context) {

	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "生活习惯",
			Name:   "lry",
		},
		Author: "lry",
		Time:   helper.GetToday(),
	}, item())

	resp.SendXML(ctx, res)
}

func item() []rss.Item {
	items := []Notification{
		// 生活习惯
		{Prefix: "生活习惯", Task: "每周五晚上，扫地拖地、刮胡子、理发、清洗脏衣服(换内裤)", Cron: "@friday"},
		{Prefix: "生活习惯", Task: "每两周周五，打飞机，晚上洗澡的时候顺便", Cron: "@2weekly"},
		{Prefix: "生活习惯", Task: "每月20号晚上，清洗洗脸毛巾、床单枕套、枕巾、浴巾", Cron: "@monthly"},
		{Prefix: "生活习惯", Task: "每两个月，换一次洗脸仪刷头", Cron: "@2monthly"},
		// renew
		{Prefix: "renew", Task: "每两周周五，换牙刷", Cron: "@2weekly"},
		{Prefix: "renew", Task: "每半年，换跑鞋、洗脸毛巾", Cron: "@6monthly"},
		{Prefix: "renew", Task: "每年，换一次内裤、床单、枕套", Cron: "@yearly"},
	}

	ret := []rss.Item{}
	for _, item := range items {
		isFriday := carbon.Now().IsFriday()
		if isFriday && item.Cron == "@friday" {
			ret = append(ret, rss.Item{
				Title: fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Time:  helper.GetToday(),
			})
		}

		// @monthly
		dayOfMonth := carbon.Now().DayOfMonth()
		if item.Cron == "@monthly" && dayOfMonth == 20 {
			ret = append(ret, rss.Item{
				Title: fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Time:  helper.GetToday(),
			})
		}

		// @yearly
		isJanuary := carbon.Now().IsJanuary()
		if item.Cron == "@yearly" && isJanuary && dayOfMonth == 1 {
			ret = append(ret, rss.Item{
				Title: fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Time:  helper.GetToday(),
			})
		}

		// @2monthly
		if item.Cron == "@2monthly" && garray.NewIntArrayFrom([]int{1, 3, 5, 7, 9, 11}).Contains(dayOfMonth) {
			ret = append(ret, rss.Item{
				Title: fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Time:  helper.GetToday(),
			})
		}

		// @6monthly
		if item.Cron == "@6monthly" && garray.NewIntArrayFrom([]int{1, 7}).Contains(dayOfMonth) {
			ret = append(ret, rss.Item{
				Title: fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Time:  helper.GetToday(),
			})
		}

		// @2weekly
		weekOfYear := carbon.Now().WeekOfYear()
		if item.Cron == "2weekly" && weekOfYear%2 != 0 && isFriday {
			ret = append(ret, rss.Item{
				Title: fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Time:  helper.GetToday(),
			})
		}
	}

	return ret
}
