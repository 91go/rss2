package life

import (
	"fmt"
	"github.com/91go/rss2/utils/helper"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
	"time"
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
		{Prefix: Day, Task: "起床，洗漱、吃饭、歇会看看feed", StartTime: Day + "7:00", Duration: "30min", TimeStub: "7h"},
		{Prefix: Day, Task: "写代码/背面试题", StartTime: Day + "7:30", Duration: "60min", TimeStub: "7h30m"},
		{Prefix: Day, Task: "跑步5km/出门上班", StartTime: Day + "8:30", TimeStub: "8h30m"},
		{Prefix: Day, Task: "test", StartTime: Day + "10:30", TimeStub: "10h30m"},
		{Prefix: Day, Task: "test2", StartTime: Day + "11:30", TimeStub: "11h30m"},
		{Prefix: Day, Task: "test3", StartTime: Day + "2:00", TimeStub: "14h"},
		{Prefix: Day, Task: "test4", StartTime: Night + "3:00", TimeStub: "15h"},
		// 蹲坑

		// NIGHT
		{Prefix: Night, Task: "吃晚饭", StartTime: Night + "7:00", Duration: "30min", TimeStub: "19h"},
		{Prefix: Night, Task: "打卡、下班", StartTime: Night + "7:30", Duration: "30min", TimeStub: "19h30m"},
		{Prefix: Night, Task: "专心跑步，9点半之前跑步1h", StartTime: Night + "8:00", Duration: "90min", TimeStub: "20h", Remark: "通常7点到8点下班，8点之前可以到家"},
		{Prefix: Night, Task: "洗澡、休息", StartTime: Night + "9:30", Duration: "30min", TimeStub: "21h30m"},
		{Prefix: Night, Task: "泡脚、写代码、看视频", StartTime: Night + "10:00", Duration: "60min", TimeStub: "22h"},
		{Prefix: Night, Task: "睡觉", StartTime: Night + "11:00", TimeStub: "23h"},
	}
)

func HabitRoutineRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "routine",
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
			title = fmt.Sprintf("[%s%s]-[从%s开始，大概%s]-%s", gtime.Date(), item.Prefix, item.StartTime, item.Duration, item.Task)
		} else {
			title = fmt.Sprintf("[%s%s]-[从%s开始]-%s", gtime.Date(), item.Prefix, item.StartTime, item.Task)
		}

		ret = append(ret, rss.Item{
			Title:       title,
			Contents:    item.Remark,
			UpdatedTime: CheckDateTime(item.TimeStub),
			ID:          rss.GenerateDateGUID("habit-routine", item.Task),
		})
	}
	return ret
}

// 与当前时间对比
func CheckDateTime(nn string) time.Time {
	// str, err := gtime.NewFromStr(gtime.Datetime()).AddStr(nn)
	str, err := gtime.NewFromTime(helper.GetToday()).AddStr(nn)
	if err != nil {
		fmt.Println(err)
		return time.Time{}
	}

	return str.Time
}
