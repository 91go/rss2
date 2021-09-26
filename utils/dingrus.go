package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/91go/rss2/core/config"

	"github.com/sirupsen/logrus"
)

type DingResponse struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type DingHook struct {
	Webhook *url.URL // 钉钉机器人的 Webhook URL
	client  *http.Client
}

const (
	DingBaseUrl = "https://oapi.dingtalk.com/robot/send?access_token="
)

func NewDingHook(webhook string, client *http.Client) (*DingHook, error) {
	wh, err := url.Parse(webhook)
	if err != nil {
		return nil, errors.New("Parse webhook to url.URL error: " + err.Error())
	}

	dh := &DingHook{Webhook: wh}
	if client != nil {
		dh.client = client
	} else {
		dh.client = &http.Client{}
	}

	return dh, err
}

func (dh *DingHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (dh *DingHook) Fire(entry *logrus.Entry) error {
	b, err := json.Marshal(entry.Data)
	if err != nil {
		return errors.New("Marshal Fields to JSON error: " + err.Error())
	}

	body := ioutil.NopCloser(bytes.NewBuffer(b))
	request := &http.Request{
		Method:     "POST",
		URL:        dh.Webhook,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Body:       body,
		Host:       dh.Webhook.Host,
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := dh.client.Do(request)
	if err != nil {
		return errors.New("Send to DingTalk error: " + err.Error())
	}
	defer response.Body.Close()

	rb, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Read DingTalk response error: " + err.Error())
	}

	dr := &DingResponse{}
	err = json.Unmarshal(rb, dr)
	if err != nil {
		return errors.New("Unmarshal DingTalk response to JSON error: " + err.Error())
	}

	if dr.ErrCode != 0 {
		return errors.New("DingTalk return error message: " + dr.ErrMsg)
	}

	return nil
}

// AssembleUrl 组装url
func AssembleUrl() string {
	//  把timestamp+"\n"+密钥当做签名字符串，使用HmacSHA256算法计算签名，然后进行Base64 encode，最后再把签名参数再进行urlEncode，得到最终的签名（需要使用UTF-8字符集）。
	token := config.GetString("dingtalk.token")
	secret := config.GetString("dingtalk.secret")

	now := time.Now().UnixMilli()
	signStr := fmt.Sprintf("%d\n%s", now, secret)

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(signStr))
	sum := hash.Sum(nil)

	encode := base64.StdEncoding.EncodeToString(sum)
	urlEncode := url.QueryEscape(encode)

	// 构建 请求 url
	return fmt.Sprintf("%s%s&timestamp=%d&sign=%s", DingBaseUrl, token, now, urlEncode)
}
