package helper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/yuin/goldmark/extension"

	"github.com/yuin/goldmark/renderer/html"

	toc "github.com/abhinav/goldmark-toc"
	"github.com/gogf/gf/os/gfile"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"

	"github.com/gogf/gf/os/gtime"
	"github.com/sirupsen/logrus"
)

// GetToday 获取今天的零点时间
func GetToday() time.Time {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t
}

// StrToTime 字符串转time.Time
func StrToTime(str, format string) time.Time {
	st, err := gtime.StrToTimeFormat(str, format)
	if err != nil {
		return time.Time{}
	}
	return st.Time
}

func TransTime(str string) time.Time {
	format, err := gtime.StrToTimeFormat(str, "Y/n/d H:i:s")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"time": str,
			"err":  err,
		}).Warn("transTime failed")
		return time.Time{}
	}
	return format.Time
}

// TrimBlank 移除HTML的空格
func TrimBlank(str string) string {
	t := strings.Replace(str, " ", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "&nbsp", "", -1)
	t = strings.Replace(t, " ", "", -1)
	return t
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// markdown转HTML
func Md2HTML(md string) string {
	if md == "" {
		return ""
	}
	var buf bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithExtensions(
			// TOC拓展
			&toc.Extender{},
			extension.GFM,
			// 删除线
			extension.Strikethrough,
			extension.TaskList,
			extension.Linkify,
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := markdown.Convert([]byte(md), &buf); err != nil {
		return ""
	}

	return buf.String()
}

// 随机字符串
func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// 获取文件MimeType
func GetContentType(filePath string) string {
	file, err := gfile.Open(filePath)
	if err != nil {
		return "audio/mpeg"
	}
	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	contentType := http.DetectContentType(buffer)
	return contentType
}

func GetAllFiles(dir string) ([]string, error) {
	dirPath, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	sep := string(os.PathSeparator)
	for _, fi := range dirPath {
		if fi.IsDir() { // 如果还是一个目录，则递归去遍历
			subFiles, err := GetAllFiles(dir + sep + fi.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			files = append(files, dir+sep+fi.Name())
		}
	}
	return files, nil
}
