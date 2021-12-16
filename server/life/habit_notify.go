package life

import (
	"fmt"
	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/os/gtime"
	"github.com/golang-module/carbon"
)

const (
	LifeHabit = "生活习惯"
	Clean     = "清洗/清洁"
	Renew     = "更换"
	ReBuy     = "复购/更新"
	FoodBuy   = "食物采购"
)

const (
	TwoDaily     = "@2daily"
	ThreeDaily   = "@3daily"
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

var CronTime = map[string]string{
	"@2daily":   "每两天",
	"@3daily":   "每三天",
	"@6daily":   "每六天",
	"@weekly":   "每周",
	"@2weekly":  "每两周",
	"@3weekly":  "每三周",
	"@monthly":  "每月",
	"@2monthly": "每两个月",
	"@3monthly": "每三个月",
	"@6monthly": "每半年",
	"@yearly":   "每年",
}

const (
	HairCut = `1. 自己理发有什么要注意的？
    1. *头发长度保持24mm，每周理发一次，打理起来很方便*
    2. *两边的头发不用太短，把鬓角推掉就可以了，其他地方不用管*，推太短了头皮露出来不好看
    3. *后面的头发，在看不到的情况下，很难处理，所以就不要弄了*，习惯习惯就好了
    4. 理完之后，站定看看齐不齐，换角度看看`
)

type Notification struct {
	Prefix string // 前缀
	Task   string // 任务
	Cron   string // 执行时间
	Remark string // 备注
}

var notifications = []Notification{
	// 生活习惯
	{Prefix: LifeHabit, Task: "每周五：刮胡子、换牙刷", Cron: Weekly},
	{Prefix: LifeHabit, Task: "每周五：理发", Cron: Weekly, Remark: helper.Md2HTML(HairCut)},
	{Prefix: LifeHabit, Task: "每两周周五：打飞机，晚上洗澡的时候顺便", Cron: TwoWeekly},
	{Prefix: LifeHabit, Task: "每两周周五：剪手指甲", Cron: TwoWeekly},
	{Prefix: LifeHabit, Task: "每两个月：换洗脸仪刷头", Cron: TwoMonthly},
	{Prefix: LifeHabit, Task: "每两个月：剪脚趾甲", Cron: TwoMonthly},
	// ???
	// todo
	{Prefix: LifeHabit, Task: "每周五：写周报，评估是否完成habit", Cron: Weekly},
	// 食物采购
	{Prefix: FoodBuy, Task: "每三天：蔬菜，买三袋", Cron: ThreeDaily, Remark: "莲藕、西芹、四季豆、西兰花、香菇、豌豆、春笋"},
	{Prefix: FoodBuy, Task: "每三天：脱脂奶，买一桶1.4L(平均每天500ml)", Cron: ThreeDaily},
	{Prefix: FoodBuy, Task: "每三天：肉类，买一袋300g(平均每天100g)", Cron: ThreeDaily},
	{Prefix: FoodBuy, Task: "每三天：苹果，买一袋(4个)", Cron: ThreeDaily},
	{Prefix: FoodBuy, Task: "每六天：燕麦，买一袋500g(平均每天100g左右)", Cron: SixDaily},
	{Prefix: FoodBuy, Task: "每六天：鸡蛋，买一盒(6个装)", Cron: SixDaily},
	// 更换
	{Prefix: Renew, Task: "每两天：换袜子、内裤", Cron: TwoDaily, Remark: ""},
	{Prefix: Renew, Task: "每周五：换速干衣(如果冬天还有速干秋裤)、睡衣睡裤，外衣外裤是否更换看需要", Cron: Weekly},
	// 清洗
	{Prefix: Clean, Task: "每周五：扫地拖地", Cron: Weekly},
	{Prefix: Clean, Task: "每周五：清洗本周脏衣服", Cron: Weekly},
	{Prefix: Clean, Task: "每月：清洗洗脸毛巾、床单枕套、枕巾、浴巾", Cron: Monthly},
	// 复购
	{Prefix: ReBuy, Task: "每2周：牙刷", Cron: TwoWeekly},
	{Prefix: ReBuy, Task: "每2月：抽纸(4包)、牙膏(黑人牙膏190g)、洗面奶(uno-130g)", Cron: TwoMonthly},
	{Prefix: ReBuy, Task: "每2月：牙线(屈臣氏50支*2)、擦镜纸(100片)", Cron: TwoMonthly, Remark: "平均每天2支/2片，所以每两个月复购一次"},
	{Prefix: ReBuy, Task: "每3月：毛巾(三利*2)", Cron: ThreeMonthly},
	{Prefix: ReBuy, Task: "每3月：湿巾(gatsby-42片*2)", Cron: ThreeMonthly, Remark: "平均每天一片"},
	{Prefix: ReBuy, Task: "每半年：跑鞋", Cron: SixMonthly},
	{Prefix: ReBuy, Task: "每半年：洗衣液(500g)", Cron: SixMonthly},
	{Prefix: ReBuy, Task: "每半年：洗发水(900g)", Cron: SixMonthly, Remark: "用洗发水代替沐浴露、洗手液、洗洁精"},
	{Prefix: ReBuy, Task: "每年：内裤(4条)", Cron: Yearly},
	{Prefix: ReBuy, Task: "每年：床笠、枕套", Cron: Yearly},
	{Prefix: ReBuy, Task: "每年：搓澡巾(单只)、鼻通(6支)", Cron: Yearly},
	{Prefix: ReBuy, Task: "每年：垃圾袋(100只)", Cron: Yearly, Remark: "平均每周两袋垃圾，一年正好用100只垃圾袋"},
	{Prefix: ReBuy, Task: "每年：新年给父母¥1000", Cron: Yearly},
}

// 用rss代替"提醒事项APP"的原因是，
func HabitNotifyRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "生活习惯notification",
		},
		Author:      "lry",
		UpdatedTime: helper.GetToday(),
	}, habitFeed())

	resp.SendXML(ctx, res)
}

func habitFeed() []rss.Item {
	ret := []rss.Item{}
	for _, item := range notifications {
		if CheckCron(item.Cron, carbon.Now()) {
			ret = append(ret, rss.Item{
				Title:       fmt.Sprintf("[%s] - [%s] - [%s] - %s", item.Prefix, gtime.Date(), CronTime[item.Cron], item.Task),
				Contents:    item.Remark,
				UpdatedTime: helper.GetToday(),
				ID:          rss.GenerateDateGUID("habit-notify", item.Task),
			})
		}
	}

	return ret
}

func CheckCron(cronTime string, carbon carbon.Carbon) bool {
	isFriday := carbon.IsFriday()
	dayOfYear := carbon.DayOfYear()
	dayOfMonth := carbon.DayOfMonth()
	weekOfYear := carbon.WeekOfYear()
	monthOfYear := carbon.MonthOfYear()
	isJanuary := carbon.IsJanuary()

	// @2daily
	if cronTime == TwoDaily && ((dayOfYear-1)%2 == 0 || dayOfYear == 1) {
		return true
	}
	// @3daily
	if cronTime == ThreeDaily && ((dayOfYear-1)%3 == 0 || dayOfYear == 1) {
		return true
	}
	// @6daily
	if cronTime == SixDaily && ((dayOfYear-1)%6 == 0 || dayOfYear == 1) {
		return true
	}
	// @weekly
	if cronTime == Weekly && isFriday {
		return true
	}
	// @2weekly
	if cronTime == TwoWeekly && weekOfYear%2 != 0 && isFriday {
		return true
	}
	// // @3weekly
	// if cronTime == ThreeWeekly && weekOfYear%3 != 0 && isFriday {
	// 	return true
	// }
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
