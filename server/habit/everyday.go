package habit

import (
	"fmt"

	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
)

const (
	Morning = "早起"
	Day     = "早上"
	Night   = "晚上"
)

type Routine struct {
	Prefix, Task, Remark, Duration, TimeStub string
}

var routines = map[string][]Routine{
	// 早起时间最好有条理，两组10min+30min
	Morning: {
		{Task: "醒来", TimeStub: "6h50m", Duration: "10min", Remark: "做3*20个提肛运动，想想当天要做的事"},
		{Task: "起床+放点提气的歌+洗漱+喝杯温水+叠被子+(处理notify)", TimeStub: "7h", Duration: "10min", Remark: ReadMarkdown("everyday", "getup.md")},
		{Task: "跑步5km，顺便看看feed/代码视频/娱乐视频", TimeStub: "7h10m", Duration: "30min", Remark: ReadMarkdown("everyday", "exercise.md")},
		{Task: "吃饭+散步", TimeStub: "7h40m", Duration: "10min", Remark: "500ml牛奶+100g花卷/饼子/燕麦"},
		{Task: "写代码/背面试题", TimeStub: "7h50m", Duration: "30min", Remark: "时间不固定，具体看通勤时间"},
		{Task: "准备上班：吃水果，穿衣服", TimeStub: "8h20m", Duration: "10min", Remark: "吃水果，每天两个苹果(500g水果)"},
	},
	Day: {
		{Task: "出门", TimeStub: "8h30m", Remark: "如果要坐地铁的话，最好避开早高峰，需要调整时间，把时间往前调"},
		{Task: "吃午饭", TimeStub: "12h", Remark: "300g蔬菜+100g肉+100g荞麦面(拇指粗)/杂粮饭(半碗)"},
	},
	Night: {
		{Task: "吃晚饭", TimeStub: "19h", Duration: "30min", Remark: "50g燕麦// 100g花卷/饼子"},
		{Task: "打卡、下班", TimeStub: "19h30m", Duration: "30min"},
		{Task: "专心跑步，9点半之前跑步1h", TimeStub: "20h", Duration: "90min", Remark: "通常7点到8点下班，8点之前可以到家"},
		{Task: "洗澡", TimeStub: "21h30m", Duration: "30min", Remark: `晚洗澡都要搓澡`},
		{Task: "泡脚，做当天代码的CR/梳理当天学到的东西/", TimeStub: "22h", Duration: "60min", Remark: "*每天泡脚*，泡脚温度40度，泡脚时间，两组，大概30min；不要饭后泡脚；"},
		{Task: "蹲坑", TimeStub: "23h", Duration: "10min", Remark: "比较习惯晚上蹲坑"},
		{Task: "睡觉", TimeStub: "23h", Remark: ReadMarkdown("everyday", "sleep.md")},
	},
}

func HabitEverydayRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "每日习惯everyday",
		},
		Author:      "lry",
		URL:         GetURL(ctx.Request),
		UpdatedTime: helper.GetToday(),
	}, routineFeed())

	resp.SendXML(ctx, res)
}

func routineFeed() []rss.Item {
	ret := []rss.Item{}

	for prefix, routine := range routines {
		for _, item := range routine {
			title := ""
			dateTime := CheckDateTime(item.TimeStub)
			formatTime := dateTime.Format("H:i")

			if item.Duration != "" {
				title = fmt.Sprintf("(从%s%s开始，预计%s)%s", prefix, formatTime, item.Duration, item.Task)
			} else {
				title = fmt.Sprintf("(从%s%s开始)%s", prefix, formatTime, item.Task)
			}

			if CheckDateTime(item.TimeStub).Before(gtime.Now()) {
				ret = append(ret, rss.Item{
					Title:       title,
					Contents:    item.Remark,
					UpdatedTime: dateTime.Time,
					ID:          rss.GenDateID("habit-routine", item.Task),
				})
			}
		}
	}
	return ret
}

// 与当前时间对比
func CheckDateTime(nn string) *gtime.Time {
	str, err := gtime.NewFromTime(helper.GetToday()).AddStr(nn)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return str
}
