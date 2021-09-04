package porn

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/91go/gofc/fctime"

	"github.com/91go/rss2/core"
)

func Test_patchVideoURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{url: "https://91porny.com/video/viewhd/483fdf11d0dac82d79e9"}, "https://91porny.com/video/view/483fdf11d0dac82d79e9"},
		{"", args{url: "https://91porny.com/video/view/483fdf11d0dac82d79e9"}, "https://91porny.com/video/view/483fdf11d0dac82d79e9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := patchVideoURL(tt.args.url); got != tt.want {
				t.Errorf("patchVideoURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPorn(t *testing.T) {
	url := "https://jiuse911.com/author/Hhonswifelonely"

	doc := core.FetchHTML(url).Text()

	fmt.Println(doc)
}

func Test_getCreateTime(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"", args{text: "2021-05-10 | 2.37万次播放"}, fctime.StrToTime("2021-05-10", "")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCreateTime(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCreateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
