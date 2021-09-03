package core

import (
	"time"

	"github.com/gorilla/feeds"
)

type Feed struct {
	URL      string
	Title    string
	Time     time.Time
	Author   string
	Contents string
	Pics     string
}

const (
	LimitItem = 4
)

func Rss(summary *Feed, items []Feed) string {
	feed := &feeds.Feed{
		Title:   summary.Title,
		Link:    &feeds.Link{Href: summary.URL},
		Author:  &feeds.Author{Name: summary.Author},
		Created: items[0].Time,
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

	atom, err := feed.ToAtom()
	if err != nil {
		return ""
	}
	return atom
}
