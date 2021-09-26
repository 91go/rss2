package rss

import (
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"

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
		logrus.WithFields(utils.Text("", errors.New("")))
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
		logrus.WithFields(utils.Text("", errors.New("rss generate failed")))
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

	m := []string{}
	for _, item := range items {
		m = append(m, item.URL, gtime.Now().TimestampStr())
	}
	// 判断key是否存在，不存在则直接set并返回
	if clt.Conn.Exists(utils.Ctx, feed.URL).Val() != 1 {
		err := clt.Conn.HSet(utils.Ctx, feed.URL, m).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
		for i, item := range items {
			item.Time = gtime.Now().Time
			items[i] = item
		}
		return rss(feed, items)
	}

	// 如果更新了，把新数据append进去，再返回
	res := checkIsUpdate(clt, feed, items)
	if len(res) != 0 {
		n := []string{}
		for _, re := range res {
			n = append(n, re, gtime.Now().TimestampStr())
		}
		clt.Conn.HSet(utils.Ctx, feed.URL, n)
	}

	// 获取更新item
	old := clt.Conn.HGetAll(utils.Ctx, feed.URL).Val()
	for i, item := range items {
		if search, ok := old[item.URL]; ok {
			item.Time = gtime.NewFromTimeStamp(gconv.Int64(search)).Time
			items[i] = item
		} else {
			fmt.Println(item.URL, "key not exist")
		}
	}
	return rss(feed, items)
}

func checkIsUpdate(clt *utils.Client, feed *Feed, items []Item) []string {
	// 通过对比相同name下的key，检查item是否更新
	old := clt.Conn.HKeys(utils.Ctx, feed.URL).Val()

	neo := []string{}
	for _, item := range items {
		neo = append(neo, item.URL)
	}
	return difference(old, neo)
}

// 比较两个[]string
func difference(slice1, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
