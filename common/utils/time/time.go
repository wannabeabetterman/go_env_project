package time

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Days(timestampFrom, timestampTo int64) int {
	var midnightUnix = func(t time.Time) int64 {
		y, m, d := t.Date()
		return time.Date(y, m, d+1, 0, 0, 0, 0, time.Local).Unix()
	}

	var days = 0
	for {
		if midnightUnix(time.Unix(timestampFrom, 0).AddDate(0, 0, days)) >= timestampTo {
			days++
			break
		}
		days++
	}
	return days
}

func BuildTime(t time.Time, years, months, days, hours, minutes int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	newBeginTime := time.Date(year+years, month+time.Month(months), day+days, hour+hours, min+minutes, sec, 0, time.Local)
	return newBeginTime
}

func BuildDate(i int) string {
	if i >= 10 {
		return strconv.Itoa(i) + ":00"
	}
	return "0" + strconv.Itoa(i) + ":00"
}

func BuildHour(i int) string {
	return strconv.Itoa(i) + ":00"
}

func BuildDayDate() (data []string) {
	for i := 0; i < 24; i++ {
		data = append(data, BuildDate(i))
	}
	return data
}

func BuildHourDate() (data []string) {
	for i := 0; i < 24; i++ {
		data = append(data, BuildHour(i))
	}
	return data
}

func BuildWeekDate() (data []string) {
	for i := 0; i < 7; i++ {
		timeStr := time.Now().AddDate(0, 0, -i)
		data = append(data, timeStr.Format(timeDateOnly))
	}
	return data
}

func BuildMonthDate() (data []string) {
	for i := 0; i <= 30; i++ {
		timeStr := time.Now().AddDate(0, 0, -i)
		data = append(data, timeStr.Format(timeDateOnly))
	}
	return data
}

func BuildYearDate() (data []string) {
	for i := 0; i <= 12; i++ {
		timeStr := time.Now().AddDate(0, -i, 0)
		data = append(data, timeStr.Format(timeDateOnlyMonth))
	}
	return data
}

func BuildDateStr(i int) string {
	if i >= 10 {
		return strconv.Itoa(i)
	}
	return "0" + strconv.Itoa(i)
}

// BuildFiveMinuteDayDate
//
//	@Description: 构建一天5分钟时段
//	@return data
func BuildFiveMinuteDayDate() (data []string) {
	total := 24 * 60 / 5
	for i := 0; i < total; i++ {
		value := i * 5
		if value >= 60 {
			data = append(data, BuildDateStr(value/60)+":"+BuildDateStr(value%60))
		} else {
			data = append(data, "00:"+BuildDateStr(value%60))
		}
	}
	return data
}

// BuildFiveMinuteDayDate
//
//	@Description: 构建一天1分钟时段
//	@return data
func BuildOneMinuteDayDate() (data []string) {
	total := 24 * 60
	for i := 0; i < total; i++ {
		value := i
		if value > 60 {
			data = append(data, BuildDateStr(value/60)+":"+BuildDateStr(value%60))
		} else {
			data = append(data, "00:"+BuildDateStr(value%60))
		}
	}
	return data
}

func GetOneMonthDate() (firstOfMonth time.Time, lastOfMonth time.Time) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth = firstOfMonth.AddDate(0, 1, -1)
	return firstOfMonth, lastOfMonth
}

func GetOneYearDate() (firstOfYear time.Time, lastOfYear time.Time) {
	now := time.Now()
	currentYear, _, _ := now.Date()
	currentLocation := now.Location()

	first := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, currentLocation)

	return first, now
}

// MyTime 自定义时间
type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.ParseInLocation(timeDateTime, timeStr, time.Local)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(timeDateTime))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(timeDateTime), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

func BuildTimeDay(beginTimeStr, endTimeStr string) (timeDay []string, isDay bool) {
	beginTime, _ := time.Parse(timeDateOnly, beginTimeStr)
	endTime, _ := time.Parse(timeDateOnly, endTimeStr)
	if beginTime.IsZero() && endTime.IsZero() {
		return timeDay, true
	}
	num := Days(beginTime.Unix(), endTime.Unix())
	if num > 365 {
		return timeDay, false
	}
	for i := 0; i < num; i++ {
		nowTime := beginTime.AddDate(0, 0, i)
		timeDay = append(timeDay, nowTime.Format(timeDateOnly))
	}
	return timeDay, true
}

func BuildLocalTime(timeStr string) time.Time {

	t, _ := time.ParseInLocation(timeDateTime, timeStr, time.Local)
	return t
}

func BuildLocalShortTime(timeStr string) time.Time {

	t, _ := time.ParseInLocation(timeDateOnly, timeStr, time.Local)
	return t
}

func BuildLocalShortTimeFormat(paramTime time.Time) string {
	t := paramTime.Local().Format(timeDateOnly)
	return t
}

func BuildLocalTimeFormat(paramTime time.Time) string {

	t := paramTime.Local().Format(timeDateTime)
	return t
}

const (
	timeDateOnly      string = "2006-01-02"
	timeDateTime      string = "2006-01-02 15:04:05"
	timeDateOnlyMonth string = "2006-01"
)
