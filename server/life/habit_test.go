package life

import (
	"github.com/golang-module/carbon"
	"testing"
	"time"
)

func TestHabitTime(t *testing.T) {
	month := carbon.Now().DayOfMonth()

	month2 := carbon.Now().Month()
	year := carbon.Now().MonthOfYear()

	// 奇数周，每两周，1
	week := carbon.Now().WeekOfYear()

	t.Log(month, month2, year, week)
}

func TestCheckCron(t *testing.T) {
	type args struct {
		cronTime string
		carbon   carbon.Carbon
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/03")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/10")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/11")}, false},

		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/03")}, false},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/10")}, true},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/17")}, false},
		{"2021/12/24", args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/24")}, true},
		{"2021/12/31", args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/31")}, false}, // 第1周
		{"2022/1/07", args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/07")}, true},    // 第2周
		{"2022/1/14", args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/14")}, false},
		{"2022/1/21", args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/21")}, true},

		// {"2021/12/03", args{cronTime: ThreeWeekly, carbon: Time2Carbon("2021/12/03")}, false}, // 第49周
		// {"2021/12/10", args{cronTime: ThreeWeekly, carbon: Time2Carbon("2021/12/10")}, false}, // true
		// {"2021/12/17", args{cronTime: ThreeWeekly, carbon: Time2Carbon("2021/12/17")}, true},
		// {"2021/12/17", args{cronTime: ThreeWeekly, carbon: Time2Carbon("2021/12/24")}, false},
		// {"2021/12/17", args{cronTime: ThreeWeekly, carbon: Time2Carbon("2021/12/31")}, false}, // 2022年第1周 // true
		// {"2022/1/07", args{cronTime: ThreeWeekly, carbon: Time2Carbon("2022/1/07")}, true}, // 第2周
		// {"2022/1/14", args{cronTime: ThreeWeekly, carbon: Time2Carbon("2022/1/14")}, false}, // true
		// {"2022/1/21", args{cronTime: ThreeWeekly, carbon: Time2Carbon("2022/1/21")}, false},

		{Monthly, args{cronTime: Monthly, carbon: Time2Carbon("2021/12/01")}, true},
		{Monthly, args{cronTime: Monthly, carbon: Time2Carbon("2021/12/02")}, false},

		{TwoMonthly, args{cronTime: TwoMonthly, carbon: Time2Carbon("2021/3/01")}, true},
		{TwoMonthly, args{cronTime: TwoMonthly, carbon: Time2Carbon("2021/5/01")}, true},
		{TwoMonthly, args{cronTime: TwoMonthly, carbon: Time2Carbon("2021/6/01")}, false},

		{ThreeMonthly, args{cronTime: ThreeMonthly, carbon: Time2Carbon("2021/4/01")}, true},
		{ThreeMonthly, args{cronTime: ThreeMonthly, carbon: Time2Carbon("2021/10/01")}, true},
		{ThreeMonthly, args{cronTime: ThreeMonthly, carbon: Time2Carbon("2021/11/01")}, false},

		{SixMonthly, args{cronTime: SixMonthly, carbon: Time2Carbon("2021/7/01")}, true},
		{SixMonthly, args{cronTime: SixMonthly, carbon: Time2Carbon("2021/8/01")}, false},

		{Yearly, args{cronTime: Yearly, carbon: Time2Carbon("2022/1/01")}, true},
		{Yearly, args{cronTime: Yearly, carbon: Time2Carbon("2022/1/02")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCron(tt.args.cronTime, tt.args.carbon); got != tt.want {
				t.Errorf("CheckCron() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Time2Carbon(str string) carbon.Carbon {
	return carbon.Time2Carbon(TransTimeN(str))
}

func TransTimeN(str string) time.Time {
	tt, _ := time.Parse("2006/1/02", str)
	return tt
}
