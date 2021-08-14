package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 获取数字区间的所有数字
func GetRange(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}

	s := make([]int, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}

// 删除字符串里的[]
func DelBrackets(slice interface{}) string {
	ios, _ := json.Marshal(slice)
	str := string(ios)
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	return str
}

// 生成指定位数的随机数
func RandNum(digit int) string {
	str := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	return str
}

func EchoStr(str interface{}) string {
	return fmt.Sprintf("%s", str)
}

func NowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 三元运算符
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// [Go如何对数组切片进行去重 - 知乎](https://zhuanlan.zhihu.com/p/86995736)
func removeDuplicateSlice(slice []string) []string {
	result := make([]string, 0, len(slice))
	temp := map[string]struct{}{}
	for _, item := range slice {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// 中午12点刷一条，下午2点40刷一条
//func GetUniqueId() string {
//
//	// 判断时间是否在早上9点到下午3点
//	now := time.Now()
//	am12 := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
//	pm2 := time.Date(now.Year(), now.Month(), now.Day(), 14, 40, 0, 0, now.Location())
//
//	if now.Equal(am12) || now.Equal(pm2) {
//		node, err := snowflake.NewNode(1)
//		if err != nil {
//			fmt.Println(err)
//			return ""
//		}
//		// Generate a snowflake ID.
//		id := node.Generate()
//
//		return id.Base64()
//	}
//
//	return ""
//}
