package life

import (
	"fmt"
	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
)

const (
	Day   = "早上"
	Night = "晚上"
)

type Routine struct {
	Prefix, Task, StartTime, Remark, Duration, TimeStub string
}

var (
	routines = []Routine{
		// DAY
		// {Prefix: Day, Task: "起床，洗漱、吃饭、歇会看看feed", StartTime: Day + "7:00", Duration: "30min", TimeStub: "7h"},
		// {Prefix: Day, Task: "写代码/背面试题", StartTime: Day + "7:30", Duration: "60min", TimeStub: "7h30m"},
		// {Prefix: Day, Task: "跑步5km/出门上班", StartTime: Day + "8:30", TimeStub: "8h30m"},

		// 早起时间最好舒展而不松弛，紧张而不紧迫
		// 有可以优化的点，随时修改
		{Prefix: Day, Task: "起床+喝杯温水+叠被子+放歌+洗漱", StartTime: Day + "7:00", TimeStub: "7h", Duration: "10min"},
		{Prefix: Day, Task: "跑步5km，顺便看看feed/代码视频/娱乐视频", StartTime: Day + "7:10", TimeStub: "7h10m", Duration: "30min"},
		{Prefix: Day, Task: "吃饭+散步", StartTime: Day + "7:40", TimeStub: "7h40m", Duration: "10min"},
		{Prefix: Day, Task: "写代码/背面试题", StartTime: Day + "7:50", TimeStub: "7h50m", Duration: "30-40min", Remark: "时间不固定，具体看通勤时间"},
		{Prefix: Day, Task: "蹲坑", StartTime: Day + "8:20", TimeStub: "8h20m", Duration: "10m"},
		{Prefix: Day, Task: "出门", StartTime: Day + "8:30", TimeStub: "8h30m", Remark: "如果要坐地铁的话，最好避开早高峰，需要调整时间，把时间往前调"},

		// NIGHT
		{Prefix: Night, Task: "吃晚饭", StartTime: Night + "7:00", TimeStub: "19h", Duration: "30min"},
		{Prefix: Night, Task: "打卡、下班", StartTime: Night + "7:30", TimeStub: "19h30m", Duration: "30min"},
		{Prefix: Night, Task: "专心跑步，9点半之前跑步1h", StartTime: Night + "8:00", TimeStub: "20h", Duration: "90min", Remark: "通常7点到8点下班，8点之前可以到家"},
		{Prefix: Night, Task: "洗澡", StartTime: Night + "9:30", TimeStub: "21h30m", Duration: "30min"},
		{Prefix: Night, Task: "泡脚，写代码", StartTime: Night + "10:00", TimeStub: "22h", Duration: "60min"},
		{Prefix: Night, Task: "睡觉", StartTime: Night + "11:00", TimeStub: "23h"},
	}
)

func HabitRoutineRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "每日routine",
		},
		Author:      "lry",
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
				Contents:    item.Remark,
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
