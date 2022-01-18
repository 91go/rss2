package habit

import (
	"fmt"
	"net/http"

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
	LifeHabit = "生活习惯"
	Clean     = "清洁"
	Renew     = "更换"
	ReBuy     = "复购"
	FoodBuy   = "食物采购"
)

const (
	TwoDaily     = "@2daily"
	ThreeDaily   = "@3daily"
	SixDaily     = "@6daily"
	Weekly       = "@weekly"
	Saturday     = "@saturday"
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
	"@saturday": "每周六",
	"@2weekly":  "每两周周六",
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
    2. *两边的头发不要太短，把鬓角推掉就可以了，其他地方不用管*，推太短了头皮露出来不好看
    3. *后面的头发，在看不到的情况下，很难处理，所以就不要弄了*，习惯习惯就好了
    4. 理完之后，站定看看齐不齐，换角度看看
	5. 先用24mm卡尺过一遍，两边和后面用12mm卡尺，用剃须刀把鬓角刮掉，其他地方不要刮青了`
)

type Notification struct {
	Task   string // 任务
	Cron   string // 执行时间
	Remark string // 备注
}

var notifications = map[string][]Notification{
	// 日常习惯
	LifeHabit: {
		{Task: "每周六：剪手指甲、刮胡子、理发", Cron: Saturday, Remark: HairCut},
		{Task: "每周六：写周报", Cron: Saturday, Remark: "评估是否完成habit"},
		{Task: "每两周：换牙刷", Cron: TwoWeekly},
		{Task: "每两周：打飞机，晚上洗澡的时候顺便", Cron: TwoWeekly},
		{Task: "每月：剪脚趾甲", Cron: Monthly},
		{Task: "每两个月：换洗脸仪刷头", Cron: TwoMonthly},
	},
	// 购买食物
	FoodBuy: {
		// {Task: "每三天：脱脂奶，买一桶1.4L(平均每天500ml)", Cron: ThreeDaily},
		{Task: "每三天：苹果，买一袋(4个)", Cron: ThreeDaily},
		{Task: "每三天：花卷/饼子，买300g(平均每天100g，早晚各50g)", Cron: ThreeDaily},
		// 近期在家购买项
		{Task: "每三天：蔬菜，买1000-1500g(平均每天300-500g)", Cron: ThreeDaily},
		{Task: "每三天：肉类，买一袋300g(平均每天100g)", Cron: ThreeDaily},
	},
	// 更换
	Renew: {
		{Task: "每两天：换袜子、内裤", Cron: TwoDaily, Remark: ""},
		{Task: "每周五：换速干衣(如果冬天还有速干秋裤)、睡衣睡裤，外衣外裤是否更换看需要", Cron: Saturday},
	},
	// 清洗
	Clean: {
		{Task: "每周六：扫地拖地", Cron: Saturday},
		{Task: "每周六：清洗本周脏衣服", Cron: Saturday},
		{Task: "每月：清洗洗脸毛巾、床单枕套、枕巾、浴巾", Cron: Monthly},
	},
	// 复购
	ReBuy: {
		{Task: "每月：牙刷(2只)", Cron: Monthly},
		{Task: "每三月：抽纸(4包)，牙膏(150-200g)，洗面奶(uno-130g)", Cron: ThreeMonthly},
		{Task: "每三月：牙线(50支*2)", Cron: ThreeMonthly, Remark: "平均每天1支，所以每三个月复购一次"},
		{Task: "每三月：擦镜纸(150-200片)", Cron: TwoMonthly, Remark: "平均每天2片，所以每三个月复购一次"},
		{Task: "每三月：湿巾(gatsby-42片*2)", Cron: ThreeMonthly, Remark: "平均每天一片"},
		{Task: "每半年：毛巾(2条)", Cron: SixMonthly},
		{Task: "每半年：洗衣液(500g)", Cron: SixMonthly},
		{Task: "每半年：洗发水(900g)", Cron: SixMonthly, Remark: "用洗发水代替沐浴露、洗手液、洗洁精"},
		{Task: "每年：跑鞋", Cron: Yearly},
		{Task: "每年：内裤(4条)", Cron: Yearly},
		{Task: "每年：床笠，枕套", Cron: Yearly},
		{Task: "每年：搓澡巾(1只)，鼻通(6支)", Cron: Yearly},
		{Task: "每年：垃圾袋(100只)", Cron: Yearly, Remark: "平均每周两袋垃圾，一年大约100只垃圾袋"},
		{Task: "每年：新年给父母¥1000", Cron: Yearly},
	},
}

// 每个月清洗足浴器

// 非周期执行项
// var notificationsOnce = map[string][]Notification{
// 	ReBuy: {
// 		// 每年12月买护手霜(100g或者30g*3)，基本上一个月一支30g装，一直用到第二年2月底；
// 		{},
// 	},
// }

// 用rss代替"提醒事项APP"的原因是，
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

	for prefix, notification := range notifications {
		for _, item := range notification {
			if CheckCron(item.Cron, carbon.Now()) {
				title := fmt.Sprintf("[%s] - [%s] - [%s] - %s", prefix, gtime.Date(), CronTime[item.Cron], item.Task)
				ret = append(ret, rss.Item{
					Title:       title,
					Contents:    html.Md2HTML(item.Remark),
					UpdatedTime: time.GetToday(),
					ID:          rss.GenDateID("habit-notify", item.Task),
				})
			}
		}
	}

	return ret
}

func CheckCron(cronTime string, cb carbon.Carbon) bool {
	isFriday := cb.IsFriday()
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
	// @6daily
	if cronTime == SixDaily && ((dayOfYear-1)%6 == 0 || dayOfYear == 1) {
		return true
	}
	// @weekly
	if cronTime == Weekly && isFriday {
		return true
	}
	// @saturday
	if cronTime == Saturday && isSaturday {
		return true
	}
	// @2weekly
	if cronTime == TwoWeekly && weekOfYear%2 != 0 && isSaturday {
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

func GetURL(r *http.Request) string {
	scheme := "http://"

	if r.TLS != nil {
		scheme = "https://"
	}
	return scheme + r.Host + r.RequestURI
}
