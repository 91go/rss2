package porn

import (
	"testing"
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
