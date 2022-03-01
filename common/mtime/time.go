package mtime

import (
	"time"
)

const (
	// TimeTemplate1 时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	TimeTemplate1 = "2006-01-02 15:04:05"
	TimeTemplate2 = "2006/01/02 15:04:05"
	TimeTemplate3 = "2006-01-02"
	TimeTemplate4 = "15:04:05"
	TimeTemplate5 = "20060102"
	TimeTemplate6 = "20060102150405"
	TimeTemplate7 = "2006010215"
	TimeTemplate8 = "200601"
)

type TimeOpts struct {
	Years  int
	Months int
	Days   int
}

// GetTime 获取当前时间的时间戳, 程序需要获取时间的模块都使用该函数
// opt 时间差, 不为0时，为对应时间后的时间戳
func GetTime(opts ...TimeOpts) int64 {
	var (
		years  = 0
		months = 0
		days   = 0
	)
	for _, opt := range opts {
		if opt.Years != 0 {
			years = opt.Years
		}
		if opt.Months != 0 {
			months = opt.Months
		}
		if opt.Days != 0 {
			days = opt.Days
		}
	}
	var timeObj = time.Now()
	timeObj = timeObj.AddDate(years, months, days)
	return timeObj.Unix()
}

// GetWeek 获取当前星期
// 0, 1, 2, 3, 4, 5, 6
func GetWeek(opt TimeOpts) int {
	var timeObj = time.Now()
	timeObj = timeObj.AddDate(opt.Years, opt.Months, opt.Days)
	return int(timeObj.Weekday())
}

// IntToString 秒级时间戳转时间
func IntToString(t int64, layout string) string {

	return time.Unix(t, 0).Format(layout)
}

func StringToInt(t string, layout string) int64 {
	stamp, _ := time.ParseInLocation(layout, t, time.Local)
	return stamp.Unix()
}

// GetUnixNanoTime 获取纳秒级时间戳 19 位
func GetUnixNanoTime() int64 {
	return time.Now().UnixNano()
}

// GetHaoTime 获取毫秒级时间戳 13 位
func GetHaoTime() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetStartDayTime 获取指定时间0点的时间戳
func GetStartDayTime(t int64) int64 {
	return StringToInt(IntToString(t, TimeTemplate5), TimeTemplate5)
}

// GetStartHourTime 获取指定时间戳的 日期起始时间
func GetStartHourTime(t int64) int64 {
	return StringToInt(IntToString(t, TimeTemplate7), TimeTemplate7)
}

// GetStartMonthTime 获取指定时间戳的 月份起始时间
func GetStartMonthTime(t int64) int64 {
	return StringToInt(IntToString(t, TimeTemplate8), TimeTemplate8)
}

// GetIntAddDateInt 获取指定时间戳时间差后的时间戳
func GetIntAddDateInt(t int64, opts TimeOpts) int64 {
	timeObj := time.Unix(t, 0)
	timeObj = timeObj.AddDate(opts.Years, opts.Months, opts.Days)
	return timeObj.Unix()
}

// GetIntToWeek 获取指定时间戳的星期
func GetIntToWeek(t int64) uint {
	timeObj := time.Unix(t, 0)
	return uint(timeObj.Weekday())
}
