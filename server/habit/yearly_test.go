package habit

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang-module/carbon"
)

const (
	TwoDaily     = "@2daily"
	ThreeDaily   = "@3daily"
	FourDaily    = "@4daily"
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

		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/04")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/11")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/12")}, false},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2021/12/18")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2022/3/12")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2022/3/19")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2022/3/26")}, true},
		{Weekly, args{cronTime: Weekly, carbon: Time2Carbon("2022/3/27")}, false},

		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/04")}, false},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/11")}, true},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/18")}, false},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2021/12/25")}, true},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/01")}, false}, // 第2周
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/08")}, true},
		{TwoWeekly, args{cronTime: TwoWeekly, carbon: Time2Carbon("2022/1/15")}, false},

		{Monthly, args{cronTime: Monthly, carbon: Time2Carbon("2021/12/01")}, true},
		{Monthly, args{cronTime: Monthly, carbon: Time2Carbon("2022/4/01")}, true},
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

func TestGetMonths(t *testing.T) {
	type args struct {
		step int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"", args{2}, []int{1, 3, 5, 7, 9, 11}},
		{"", args{3}, []int{1, 4, 7, 10}},
		{"", args{6}, []int{1, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetMonths(tt.args.step); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetMonths() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestCarbonDayOfWeek(t *testing.T) {
	isFriday := Time2Carbon("2022/3/25").IsFriday()
	fmt.Println(isFriday)
	isSaturday := Time2Carbon("2022/3/12").IsSaturday()
	fmt.Println(isSaturday)
}

func TestExtractTimeNumber(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{"", args{Weekly}, false, 1},
		{"", args{TwoWeekly}, true, 2},
		{"", args{ThreeWeekly}, true, 3},
		{"", args{TwoMonthly}, true, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ExtractTimeNumber(tt.args.t)
			if got != tt.want {
				t.Errorf("ExtractTimeNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractTimeNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
