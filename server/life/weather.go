package life

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/91go/gofc"
	"github.com/91go/rss2/core"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
)

const (
	base         = "https://tianqi.moji.com/weather/china/"
	WeatherLimit = 4
)

type Warm struct {
	Code string
	City string
	URL  string
	Ctx  string
	Time time.Time
}

// Weather site data
type Weather struct {
	City     string
	Temp     string
	Weather  string
	Air      string
	Humidity string
	Wind     string
	Limit    string
	Note     string
}

// WeatherRss 天气feed
func WeatherRss(ctx *gin.Context) {
	city := ctx.Query("city")

	warm := crawl(city)

	feed := &feeds.Feed{
		Title:       warm.City,
		Link:        &feeds.Link{Href: warm.URL},
		Description: warm.City,
		Author:      &feeds.Author{Name: "", Email: ""},
		Created:     warm.Time,
		Updated:     warm.Time,
	}

	feed.Add(&feeds.Item{
		Title:       warm.City,
		Link:        &feeds.Link{Href: warm.URL},
		Description: warm.City,
		Author:      &feeds.Author{Name: "", Email: ""},
		Content:     warm.Ctx,
		Created:     warm.Time,
		Updated:     warm.Time,
	})

	res, err := feed.ToAtom()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"city": city,
			"err":  err,
		}).Warn("generate weather feed failed")
	}

	core.SendXML(ctx, res)
}

func crawl(city string) Warm {
	parts := make(map[string]interface{})
	weather := getWeather(city)
	parts["weather"] = weather
	html := gofc.GenerateHTML(HTML, parts)

	return Warm{
		Code: city,
		City: weather.City,
		URL:  base + city,
		Ctx:  html,
		Time: time.Now(),
	}
}

// 聚合所有数据
// func getParts() map[string]interface{} {
// 	wrapMap := map[string]func() interface{}{
// 		//"one": func() interface{} { return GetONE() },
// 		//"poem": func() interface{} { return GetPoem() },
// 	}
//
// 	wg := sync.WaitGroup{}
// 	parts := map[string]interface{}{}
// 	for name, getPart := range wrapMap {
// 		wg.Add(1)
// 		go func(key string, fn func() interface{}) {
// 			defer wg.Done()
// 			parts[key] = fn()
// 		}(name, getPart)
// 	}
// 	wg.Wait()
// 	return parts
// }

const HTML = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>每日一暖, 温情一生</title>
</head>
<body>
  <div style="max-width: 375px; margin: 20px auto;color:#444; font-size: 16px;">
    <h3 style="text-align: center">{{weather.City}}</h3>
    <br>
    <div style="padding: 0;width: 100%;">
      <div><span style="color: #6e6e6e">天气：</span>{{weather.Weather}}</div>
      <div><span style="color: #6e6e6e">温度：</span>{{weather.Temp}}</div>
      <div><span style="color: #6e6e6e">湿度：</span>{{weather.Humidity}}</div>
      <div><span style="color: #6e6e6e">风向：</span>{{weather.Wind}}</div>
      <div><span style="color: #6e6e6e">空气：</span>{{weather.Air}}</div>
      <div><span style="color: #6e6e6e">限行：</span>{{weather.Limit}}</div>
      <div><span style="color: #6e6e6e">提示：</span>{{weather.Note}}</div>
    </div>
  </div>
  <br><br>
</body>
</html>
`

// GetWeather data
func getWeather(local string) Weather {
	url := "https://tianqi.moji.com/weather/china/" + local
	doc := core.FetchHTML(url)
	wrap := doc.Find(".wea_info .left")
	humidityDesc := strings.Split(wrap.Find(".wea_about span").Text(), " ")
	humidity := "未知"
	if len(humidityDesc) >= 2 {
		humidity = humidityDesc[1]
	}

	limitDesc := []rune(wrap.Find(".wea_about b").Text())
	limit := ""
	if len(limitDesc) <= WeatherLimit {
		limit = string(limitDesc)
	} else {
		limit = string(limitDesc[WeatherLimit:])
	}
	return Weather{
		City:     doc.Find("#search .search_default em").Text(),
		Temp:     wrap.Find(".wea_weather em").Text() + "°",
		Weather:  wrap.Find(".wea_weather b").Text(),
		Air:      wrap.Find(" .wea_alert em").Text(),
		Humidity: humidity,
		Wind:     wrap.Find(".wea_about em").Text(),
		Limit:    limit,
		Note:     strings.ReplaceAll(wrap.Find(".wea_tips em").Text(), "。", ""),
	}
}
