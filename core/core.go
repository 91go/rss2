package core

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gorilla/feeds"
	"time"
)

type Feed struct {
	Url      string
	Title    string
	Time     time.Time
	Author   string
	Contents string
	Pics     string
}

func Rss(summary Feed, items []Feed) string {
	feed := &feeds.Feed{
		Title:   summary.Title,
		Link:    &feeds.Link{Href: summary.Url},
		Author:  &feeds.Author{Name: summary.Author},
		Created: items[0].Time,
	}
	for _, value := range items {
		feed.Add(&feeds.Item{
			Title:       value.Title,
			Link:        &feeds.Link{Href: value.Url},
			Description: value.Contents,
			Author:      &feeds.Author{Name: value.Author},
			Created:     value.Time,
		})
	}

	atom, err := feed.ToAtom()
	if err != nil {
		glog.Error(err)
	}
	return atom
}
