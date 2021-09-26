package rss

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"

	"github.com/91go/rss2/utils"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/feeds"
)

// Feed 通用Feed
type Feed struct {
	URL string
	Title
	Time     time.Time
	Author   string
	Contents string
}

type Title struct {
	Prefix string
	Name   string
}

//
type Item struct {
	URL      string
	Title    string
	Time     time.Time
	Author   string
	Contents string
}

const (
	LimitItem = 3
)

// Rss 输出rss
func Rss(fe *Feed, items []Item) string {
	if len(items) == 0 {
		logrus.WithFields(utils.Fields("", errors.New("")))
		return ""
	}

	if fe.Time.IsZero() {
		return feedWithoutTime(fe, items)
	}

	return rss(fe, items)
}

func rss(fe *Feed, items []Item) string {
	feed := feeds.Feed{
		Title:   feedTitle(fe.Title),
		Link:    &feeds.Link{Href: fe.URL},
		Author:  &feeds.Author{Name: fe.Author},
		Created: fe.Time,
	}

	for _, value := range items {
		feed.Add(&feeds.Item{
			Title:       value.Title,
			Link:        &feeds.Link{Href: value.URL},
			Description: value.Contents,
			Author:      &feeds.Author{Name: value.Author},
			Created:     value.Time,
		})
	}

	atom, err := feed.ToRss()
	if err != nil {
		return ""
	}
	return atom
}

func feedTitle(tt Title) string {
	if tt.Name == "" {
		return tt.Prefix
	}
	return fmt.Sprintf("%s-%s", tt.Prefix, tt.Name)
}

// 处理没有提供更新时间的feed
func feedWithoutTime(feed *Feed, items []Item) string {
	clt := utils.NewClient(utils.Conn())

	res := StructSliceToMapSlice(items)
	g.Dump(res)

	any := gconv.SliceAny(items)
	g.Dump(any)

	// 判断key是否存在，不存在则直接set并返回
	if clt.Conn.Exists(utils.Ctx, feed.URL).Val() != 1 {
		// todo
		err := clt.Conn.LPushX(utils.Ctx, feed.URL, any).Err()

		if err != nil {
			fmt.Println(err)
			return ""
		}
		return rss(feed, items)
	}

	// 未更新
	if checkIsUpdate(clt, feed, items) {
		return rss(feed, items)
	}
	// 获取更新item
	// 有更新，把新数据append进去，再返回
	clt.Conn.LPushX(utils.Ctx, feed.URL)
	clt.Conn.LRange(utils.Ctx, feed.URL, 0, -1).Val()
	// todo []string转[]Item

	return ""
}

func checkIsUpdate(clt *utils.Client, feed *Feed, items []Item) bool {
	// 通过对比相同name下的key，检查item是否更新
	old := clt.Conn.LLen(utils.Ctx, feed.URL).Val()

	return len(items) == int(old)
}

// StructSliceToMapSlice : struct切片转为map切片
func StructSliceToMapSlice(source interface{}) []map[string]interface{} {
	// 判断，interface转为[]interface{}
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Slice {
		panic("ERROR: Unknown type, slice expected.")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	// 转换之后的结果变量
	res := make([]map[string]interface{}, 0)

	// 通过遍历，每次迭代将struct转为map
	for _, elem := range ret {
		data := make(map[string]interface{})
		objT := reflect.TypeOf(elem)
		objV := reflect.ValueOf(elem)
		for i := 0; i < objT.NumField(); i++ {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
		res = append(res, data)
	}

	return res
}
