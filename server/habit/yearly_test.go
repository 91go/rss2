package habit

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-module/carbon"
)

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
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2021/12/03")}, true},
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2021/12/04")}, false},
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2021/12/05")}, true},
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2021/12/06")}, false},
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2021/12/30")}, false},
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2021/12/31")}, true}, // 转过年有问题
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2022/1/01")}, true},
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2021/1/02")}, false},
		{TwoDaily, args{cronTime: TwoDaily, carbon: Time2Carbon("2021/1/03")}, true},

		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/12/03")}, true},
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/12/04")}, false},
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/12/05")}, false},
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/12/06")}, true},
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/12/30")}, true},
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/12/31")}, false}, // 转过年有问题
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2022/1/01")}, true},
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/1/02")}, false},
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/1/03")}, false},
		{ThreeDaily, args{cronTime: ThreeDaily, carbon: Time2Carbon("2021/1/04")}, true},

		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/03")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/10")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/11")}, false},

		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/04")}, false},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/11")}, true},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/18")}, false},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/25")}, true},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/01")}, false}, // 第2周
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/08")}, true},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/15")}, false},

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

func TestTime(t *testing.T) {
	t.Run("", func(t *testing.T) {
		for i := 1; i <= 1000; i++ {
			kk := (i - 1) % 2

			if kk == 0 {
				fmt.Println(i, ":", "value:", kk, "true")
			}
		}
	})

	t.Run("", func(t *testing.T) {
		for i := 1; i <= 1000; i++ {
			kk := (i - 1) % 3

			if kk == 0 {
				fmt.Println(i, ":", "value:", kk, "true")
			}
		}
	})
}
