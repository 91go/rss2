package code

import (
	"fmt"
	"path/filepath"

	"github.com/91go/rss2/utils/helper/time"

	"github.com/91go/rss2/utils/gq"
	"github.com/91go/rss2/utils/resp"
	"github.com/91go/rss2/utils/rss"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
)

const OneTabBaseURL = "https://www.one-tab.com/page/"

func OneTabSharedRSS(ctx *gin.Context) {
	page := ctx.Param("page")
	url := fmt.Sprintf("%s%s", OneTabBaseURL, page)

	res := rss.Rss(&rss.Feed{
		URL: url,
		Title: rss.Title{
			Prefix: "onetab",
			Name:   page,
		},
		UpdatedTime: time.GetToday(),
	}, sharedList(url))

	resp.SendXML(ctx, res)
}

func sharedList(url string) []rss.Item {
	doc := gq.FetchHTML(url)
	ret := []rss.Item{}
	doc.Find("body").Find("div").Slice(7, -1).Each(func(i int, sel *query.Selection) {
		title := sel.Find("a").Text()
		if url, exists := sel.Find("a").Attr("href"); exists {
			ret = append(ret, rss.Item{
				Title:       title,
				URL:         url,
				ID:          rss.GenFixedID("onetab-shared", url),
				UpdatedTime: time.GetToday(),
			})
		}
	})
	return ret
}

func OneTabTXTRSS(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		URL: "",
		Title: rss.Title{
			Prefix: "onetab",
			Name:   "txt",
		},
		UpdatedTime: time.GetToday(),
	}, txtList())

	resp.SendXML(ctx, res)
}

func txtList() (ret []rss.Item) {
	abs, err := filepath.Abs("./public/txt/onetab.txt")
	if err != nil {
		return nil
	}

	err = gfile.ReadLines(abs, func(text string) error {
		if text != "" {
			explode := gstr.Explode("|", text)
			url, title := explode[0], explode[1]
			ret = append(ret, rss.Item{
				Title:       title,
				URL:         url,
				UpdatedTime: time.GetToday(),
				ID:          rss.GenFixedID("onetab-txt", url),
			})
		}
		return nil
	})

	if err != nil {
		return ret
	}

	return ret
}
