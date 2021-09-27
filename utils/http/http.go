package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// RequestGet GET请求
func RequestGet(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		// 写入zap之类的数据
		fmt.Printf("Error fetching url: %s err: %v", url, err)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		fmt.Printf("Error Reading failed: %s", url)
		return nil
	}
	return body
}
