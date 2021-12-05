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

const (
	LifeHabit = "生活习惯"
	Renew     = "repurchase/renew"
)

const (
	Weekly       = "@weekly"
	TwoWeekly    = "@2weekly"
	Monthly      = "@monthly"
	TwoMonthly   = "@2monthly"
	ThreeMonthly = "@3monthly"
	SixMonthly   = "@6monthly"
	Yearly       = "@yearly"
)

type Notification struct {
	Prefix string // 前缀
	Task   string // 任务
	Cron   string // 执行时间
	Remark string // 备注
}

// 用rss代替"提醒事项APP"的原因是，
func HabitRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "生活习惯",
			Name:   "lry",
		},
		Author:      "lry",
		UpdatedTime: helper.GetToday(),
	}, item())

	resp.SendXML(ctx, res)
}

func item() []rss.Item {
	items := []Notification{
		// 生活习惯
		{Prefix: LifeHabit, Task: "每周五晚上，扫地拖地、刮胡子、理发、清洗脏衣服(换内裤)、换牙刷", Cron: Weekly},
		{Prefix: LifeHabit, Task: "每两周周五，打飞机，晚上洗澡的时候顺便", Cron: TwoWeekly},
		{Prefix: LifeHabit, Task: "每月20号晚上，清洗洗脸毛巾、床单枕套、枕巾、浴巾", Cron: Monthly},
		{Prefix: LifeHabit, Task: "每两个月，换一次洗脸仪刷头", Cron: TwoMonthly},
		// renew复购
		{Prefix: Renew, Task: "每2周：牙刷", Cron: TwoWeekly},
		{Prefix: Renew, Task: "每2月：抽纸(4包)、牙膏(黑人190g)、洗面奶(uno-130g)", Cron: TwoMonthly},
		{Prefix: Renew, Task: "每2月：牙线(屈臣氏50支*2)、擦镜纸(100片)", Remark: "平均每天2支，所以每两个月复购一次", Cron: TwoMonthly},
		{Prefix: Renew, Task: "每3月：毛巾(三利*2)", Cron: ThreeMonthly},
		{Prefix: Renew, Task: "每3月：湿巾(42片*2)", Remark: "平均每天一片", Cron: ThreeMonthly},
		{Prefix: Renew, Task: "每半年：跑鞋、洗衣液(500g)", Cron: SixMonthly},
		{Prefix: Renew, Task: "每半年：洗发水(900g)", Cron: SixMonthly, Remark: "用洗发水代替沐浴露、洗手液、洗洁精"},
		{Prefix: Renew, Task: "每年：内裤、床单、枕套", Cron: Yearly},
		{Prefix: Renew, Task: "每年：搓澡巾(单只)、鼻通(6支)", Cron: Yearly},
		{Prefix: Renew, Task: "每年：垃圾袋(100只)", Remark: "平均每周两袋垃圾，一年正好用100只垃圾袋", Cron: Yearly},
	}

	ret := []rss.Item{}
	for _, item := range items {
		isFriday := carbon.Now().IsFriday()
		dayOfMonth := carbon.Now().DayOfMonth()
		weekOfYear := carbon.Now().WeekOfYear()
		monthOfYear := carbon.Now().MonthOfYear()
		isJanuary := carbon.Now().IsJanuary()

		// @weekly
		if isFriday && item.Cron == Weekly {
			ret = append(ret, rss.Item{
				Title:    fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Contents: item.Remark,
				Time:     helper.GetToday(),
			})
		}

		// @2weekly
		if item.Cron == TwoWeekly && weekOfYear%2 != 0 && isFriday {
			ret = append(ret, rss.Item{
				Title:    fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Contents: item.Remark,
				Time:     helper.GetToday(),
			})
		}

		// @3weekly
		if item.Cron == TwoWeekly && weekOfYear%3 != 0 && isFriday {
			ret = append(ret, rss.Item{
				Title:    fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Contents: item.Remark,
				Time:     helper.GetToday(),
			})
		}

		// @monthly
		if item.Cron == Monthly && dayOfMonth == 20 {
			ret = append(ret, rss.Item{
				Title:    fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Contents: item.Remark,
				Time:     helper.GetToday(),
			})
		}

		// @2monthly
		if item.Cron == TwoMonthly && garray.NewIntArrayFrom([]int{1, 3, 5, 7, 9, 11}).Contains(monthOfYear) && dayOfMonth == 1 {
			ret = append(ret, rss.Item{
				Title:    fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Contents: item.Remark,
				Time:     helper.GetToday(),
			})
		}

		// @6monthly
		if item.Cron == SixMonthly && garray.NewIntArrayFrom([]int{1, 7}).Contains(monthOfYear) && dayOfMonth == 1 {
			ret = append(ret, rss.Item{
				Title:    fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Contents: item.Remark,
				Time:     helper.GetToday(),
			})
		}

		// @yearly
		if item.Cron == Yearly && isJanuary && dayOfMonth == 1 {
			ret = append(ret, rss.Item{
				Title:    fmt.Sprintf("[%s] - %s", item.Prefix, item.Task),
				Contents: item.Remark,
				Time:     helper.GetToday(),
			})
		}

	}

	return ret
}
