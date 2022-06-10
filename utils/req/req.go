package req

import (
	"fmt"
	"io"
	"net/http"
	"rss2/utils/log"

	"github.com/gogf/gf/net/ghttp"
	"github.com/sirupsen/logrus"
)

func Get(url string) (string, error) {
	resp, err := ghttp.NewClient().Get(url)
	if err != nil {
		logrus.WithFields(log.Text(url, nil)).Error("http request failed")
		return "", err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			logrus.WithFields(log.Text(url, nil)).Error("http close failed")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request error: %s", url)
	}
	bytes, _ := io.ReadAll(resp.Body)

	return string(bytes), nil
}
