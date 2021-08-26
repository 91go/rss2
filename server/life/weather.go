package life

import (
	"github.com/91go/gofc"
	"github.com/91go/rss2/utils"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gorilla/feeds"
	"log"
	"strings"
	"time"
)

var (
	base = "https://tianqi.moji.com/weather/china/"
)

type Warm struct {
	Code string
	City string
	Url  string
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

// One site info
type One struct {
	Date     string
	ImgURL   string
	Sentence string
}

// English info
type English struct {
	ImgURL   string
	Sentence string
}

// Poem info
type Poem struct {
	Title   string   `json:"title"`
	Dynasty string   `json:"dynasty"`
	Author  string   `json:"author"`
	Content []string `json:"content"`
}

// PoemRes response data
type PoemRes struct {
	Status string `json:"status"`
	Data   struct {
		Origin Poem `json:"origin"`
	} `json:"data"`
}

// Wallpaper data
type Wallpaper struct {
	Title  string
	ImgURL string
}

// Trivia info
type Trivia struct {
	ImgURL      string
	Description string
}

// User for receive email
type User struct {
	Email string `json:"email"`
	Local string `json:"local"`
}

func WeatherRss(request *ghttp.Request) {

	city := request.GetString("city")

	warm := crawl(city)

	feed := &feeds.Feed{
		Title:       warm.City,
		Link:        &feeds.Link{Href: warm.Url},
		Description: warm.City,
		Author:      &feeds.Author{Name: "", Email: ""},
		Created:     warm.Time,
		Updated:     warm.Time,
	}

	feed.Add(&feeds.Item{
		Title:       warm.City,
		Link:        &feeds.Link{Href: warm.Url},
		Description: warm.City,
		Author:      &feeds.Author{Name: "", Email: ""},
		Content:     warm.Ctx,
		Created:     warm.Time,
		Updated:     warm.Time,
	})

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}

	err = request.Response.WriteXmlExit(atom)
	if err != nil {
		return
	}
}

func crawl(city string) Warm {

	//parts := getParts()
	parts := make(map[string]interface{})
	weather := GetWeather(city)
	parts["weather"] = weather
	html := gofc.GenerateHTML(HTML, parts)

	return Warm{
		Code: city,
		City: weather.City,
		Url:  base + city,
		Ctx:  html,
		Time: time.Now(),
	}
}

// 聚合所有数据
//func getParts() map[string]interface{} {
//	wrapMap := map[string]func() interface{}{
//		//"one": func() interface{} { return GetONE() },
//		//"poem": func() interface{} { return GetPoem() },
//	}
//
//	wg := sync.WaitGroup{}
//	parts := map[string]interface{}{}
//	for name, getPart := range wrapMap {
//		wg.Add(1)
//		go func(key string, fn func() interface{}) {
//			defer wg.Done()
//			parts[key] = fn()
//		}(name, getPart)
//	}
//	wg.Wait()
//	return parts
//}

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
func GetWeather(local string) Weather {
	url := "https://tianqi.moji.com/weather/china/" + local
	doc := utils.FetchHTML(url)
	wrap := doc.Find(".wea_info .left")
	humidityDesc := strings.Split(wrap.Find(".wea_about span").Text(), " ")
	humidity := "未知"
	if len(humidityDesc) >= 2 {
		humidity = humidityDesc[1]
	}

	limitDesc := ([]rune)(wrap.Find(".wea_about b").Text())
	limit := ""
	if len(limitDesc) <= 4 {
		limit = string(limitDesc)
	} else {
		limit = string(limitDesc[4:])
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

// GetONE data
//func GetONE() One {
//	url := "http://wufazhuce.com/"
//	doc := utils.FetchHTML(url)
//	wrap := doc.Find(".fp-one .carousel .item.active")
//	day := wrap.Find(".dom").Text()
//	monthYear := wrap.Find(".may").Text()
//	imgURL, _ := wrap.Find(".fp-one-imagen").Attr("src")
//	return One{
//		ImgURL:   imgURL,
//		Date:     fmt.Sprintf("%s %s", day, monthYear),
//		Sentence: wrap.Find(".fp-one-cita a").Text(),
//	}
//}

// GetEnglish data
//func GetEnglish() English {
//	url := "http://dict.eudic.net/home/dailysentence"
//	doc := utils.FetchHTML(url)
//	wrap := doc.Find(".containter .head-img")
//	imgURL, _ := wrap.Find(".himg").Attr("src")
//	return English{
//		ImgURL:   imgURL,
//		Sentence: wrap.Find(".sentence .sect_en").Text(),
//	}
//}

// GetPoem data
//func GetPoem() Poem {
//	url := "https://v2.jinrishici.com/one.json"
//
//	//buf := new(bytes.Buffer)
//	//res := utils.Fetch(url)
//	//buf.ReadFrom(res)
//	//resByte := buf.Bytes()
//	resByte := utils.RequestGet(url)
//
//	var resJSON PoemRes
//	err := json.Unmarshal(resByte, &resJSON)
//	if err != nil {
//		log.Fatalf("Fetch json from %s error: %s", url, err)
//	}
//
//	status := resJSON.Status
//	if status != "success" {
//		log.Fatalf("Get poem status %s, res: %s", status, resJSON)
//	}
//	return resJSON.Data.Origin
//}
