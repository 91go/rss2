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
	Prefix, Task, StartTime, Remark, Duration, TimeStub string
}

const (
	GetUp = `1. 起床和起立时，动作要慢；避免"直立性低血压"；对身体损伤很大； 
	2. 每天早上叠被子，清扫床上异物；`
	Sleep = `1. *不熬夜，早睡早起，保持良好的睡眠* 
	2. *睡觉前把手机拿远*，从而保证睡觉前和早上起床不玩手机(用音箱做闹铃)`
	MorningExercise = `1. 晨跑前的睡眠要好，睡足8h，在睡眠和健身之间要选择睡眠；
	2. 晨跑时间小于40min或者距离小于5km，可以跑完再吃早饭；如果大量运动，则应该饭后歇至少1h再运动；
	3. 晨练不会导致低血糖；
	4. 晨练前一定要补水；
	5. 晨练一定要注意热身，因为睡了一晚之后，动态平衡能力会短暂下降，防止扭伤；
	6. 晨跑要由慢到快，让心肺功能慢慢提高；`
)

var (
	routines = []Routine{
		// Morning
		// 早起时间最好有条理，两组10min+30min
		{Prefix: Morning, Task: "醒来"},
		{Prefix: Morning, Task: "起床+放点提气的歌+洗漱+喝杯温水+叠被子+(处理notify)", StartTime: Day + "7:00", TimeStub: "7h", Duration: "10min", Remark: GetUp},
		{Prefix: Morning, Task: "跑步5km，顺便看看feed/代码视频/娱乐视频", StartTime: Day + "7:10", TimeStub: "7h10m", Duration: "30min", Remark: MorningExercise},
		{Prefix: Morning, Task: "吃饭+散步", StartTime: Day + "7:40", TimeStub: "7h40m", Duration: "10min", Remark: "500ml牛奶+100g花卷/饼子/燕麦"},
		{Prefix: Morning, Task: "写代码/背面试题", StartTime: Day + "7:50", TimeStub: "7h50m", Duration: "30min", Remark: "时间不固定，具体看通勤时间"},
		{Prefix: Morning, Task: "准备上班：吃水果，穿衣服", StartTime: Day + "8:20", TimeStub: "8h20m", Duration: "10min", Remark: "吃水果，每天两个苹果(500g水果)"},

		// DAY
		{Prefix: Day, Task: "出门", StartTime: Day + "8:30", TimeStub: "8h30m", Remark: "如果要坐地铁的话，最好避开早高峰，需要调整时间，把时间往前调"},
		{Prefix: Day, Task: "吃午饭", StartTime: Day + "12:00", TimeStub: "12h", Remark: "300g蔬菜+100g肉+100g荞麦面(拇指粗)/杂粮饭(半碗)"},

		// NIGHT
		{Prefix: Night, Task: "吃晚饭", StartTime: Night + "7:00", TimeStub: "19h", Duration: "30min", Remark: "50g燕麦// 100g花卷/饼子"},
		{Prefix: Night, Task: "打卡、下班", StartTime: Night + "7:30", TimeStub: "19h30m", Duration: "30min"},
		{Prefix: Night, Task: "专心跑步，9点半之前跑步1h", StartTime: Night + "8:00", TimeStub: "20h", Duration: "90min", Remark: "通常7点到8点下班，8点之前可以到家"},
		{Prefix: Night, Task: "洗澡", StartTime: Night + "9:30", TimeStub: "21h30m", Duration: "30min", Remark: `晚洗澡都要搓澡`},
		{Prefix: Night, Task: "泡脚，写代码", StartTime: Night + "10:00", TimeStub: "22h", Duration: "60min", Remark: "*每天泡脚*，泡脚温度40度，泡脚时间，两组，大概30min；不要饭后泡脚；"},
		// {Prefix: Night, Task: "蹲坑", StartTime: Night + "11:00", TimeStub: "23h", Duration: "10min", Remark: "比较习惯晚上蹲坑"},
		{Prefix: Night, Task: "睡觉", StartTime: Night + "11:00", TimeStub: "23h", Remark: Sleep},
	}
)

func HabitRoutineRss(ctx *gin.Context) {

	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "每日routine",
		},
		Author:      "lry",
		URL:         GetURL(ctx.Request),
		UpdatedTime: helper.GetToday(),
	}, routineFeed())

	resp.SendXML(ctx, res)
}

func routineFeed() []rss.Item {
	ret := []rss.Item{}

	for _, item := range routines {
		title := ""
		if item.Duration != "" {
			title = fmt.Sprintf("(从%s开始，预计%s)%s", item.StartTime, item.Duration, item.Task)
		} else {
			title = fmt.Sprintf("(从%s开始)%s", item.StartTime, item.Task)
		}

		if CheckDateTime(item.TimeStub).Before(gtime.Now()) {
			ret = append(ret, rss.Item{
				Title:       title,
				Contents:    helper.Md2HTML(item.Remark),
				UpdatedTime: CheckDateTime(item.TimeStub).Time,
				ID:          rss.GenerateDateGUID("habit-routine", item.Task),
			})
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
