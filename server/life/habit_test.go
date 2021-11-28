package life

import (
	"testing"

	"github.com/golang-module/carbon"
)

func TestHabitTime(t *testing.T) {
	month := carbon.Now().DayOfMonth()

	month2 := carbon.Now().Month()
	year := carbon.Now().MonthOfYear()

	// 奇数周，每两周，1
	week := carbon.Now().WeekOfYear()

	t.Log(month, month2, year, week)
}