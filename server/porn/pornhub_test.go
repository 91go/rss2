package porn

import (
	"fmt"
	"github.com/gogf/gf/text/gstr"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestFetch(t *testing.T) {

	client := resty.New()
	resp, err := client.R().EnableTrace().Get("https://cn.pornhub.com/model/mai-chen/videos?o=mr")
	if err != nil {
		return
	}
	fmt.Println(resp)
}

func TestStr(t *testing.T) {
	str := "https://www.pornhub.com/view_video.php?viewkey=ph619e0cecb5b89"
	subStr := gstr.SubStr(str, gstr.Pos(str, "=")+1)

	rightStr := gstr.TrimRightStr(str, "=")
	fmt.Println(rightStr, subStr)
}
