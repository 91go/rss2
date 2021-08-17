package utils

import (
	"strconv"
	"time"
)

// 毫秒转time.Time
func MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))

	//fmt.Println(tm.Format("2006-02-01 15:04:05.000"))
	return tm, nil
}

// 字符串转time.Time
func TransTime(str string) time.Time {
	local, _ := time.LoadLocation("Local")
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", str, local)
	return tt
}

func TransTime2(str string) time.Time {
	//local, _ := time.LoadLocation("Local")
	tt, _ := time.Parse("2006/01/02 15:04:05", str)
	return tt
}
