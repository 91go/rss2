package porn

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/91go/rss2/utils"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	url := "https://zzxcx.netzijin.cn/2021/07/20210725144915628-285x285.jpg"
	res := strings.Replace(url, "285x285", "scaled", -1)
	assert.Equal(t, "https://zzxcx.netzijin.cn/2021/07/20210725144915628-scaled.jpg", res)
}

func Test_sanitizeTime(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"", args{url: "https://youpai.netzijin.cn/2021/08/20210829140910717.jpg"}, utils.StrToTime("20210829", "Ymd")},
		{"", args{url: "https://youpai.netzijin.cn/2021/08/20210809144240113-scaled.jpg"}, utils.StrToTime("20210809", "Ymd")},
		{"", args{url: "https://youpai.netzijin.cn/2021/08/20210809142257765-scaled.jpg"}, utils.StrToTime("20210809", "Ymd")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeTime(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sanitizeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
