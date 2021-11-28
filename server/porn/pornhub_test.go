package porn

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestFetch(t *testing.T) {
	// html := gq.FetchHTML("https://cn.pornhub.com/model/mai-chen/videos?o=mr")
	// fmt.Println(html)

	client := resty.New()
	resp, err := client.R().EnableTrace().Get("https://cn.pornhub.com/model/mai-chen/videos?o=mr")
	if err != nil {
		return
	}
	fmt.Println(resp)
}
