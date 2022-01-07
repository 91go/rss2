package html

import (
	"fmt"
	"github.com/gogf/gf/text/gstr"
	"reflect"
	"strings"
)

// 生成html
func GenerateHTML(html string, datas map[string]interface{}) string {
	for key, data := range datas {
		rDataKey := reflect.TypeOf(data)
		rDataVal := reflect.ValueOf(data)
		fieldNum := rDataKey.NumField()
		for i := 0; i < fieldNum; i++ {
			fName := rDataKey.Field(i).Name
			rValue := rDataVal.Field(i)

			var fValue string
			switch rValue.Interface().(type) {
			case string:
				fValue = rValue.String()
			case []string:
				fValue = strings.Join(rValue.Interface().([]string), "<br>")
			}

			mark := fmt.Sprintf("{{%s.%s}}", key, fName)
			html = strings.ReplaceAll(html, mark, fValue)
		}
	}
	return html
}

// DealContents 根据文件类型，判断是否返回iframe
func GetIframe(filetype, tfUrl string) string {
	if gstr.Contains(filetype, "video") {
		return fmt.Sprintf(`<iframe src="%q" frameborder="0" width="640" height="390" scrolling="no" frameborder="0" border="0" framespacing="0" allowfullscreen></iframe><br><br>`, tfUrl)
	}
	return ""
}
