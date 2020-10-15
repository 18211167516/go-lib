package main

import (
	"fmt"
	"time"
)

func main() {
	testTimeInit()
	testTimeFormat()
	testTimeTimestamp()
	testTimeFuncDate()
	testTimeOperate()
}

func testTimeInit() {
	// func Now() Time 返回当前本地时间
	fmt.Println(time.Now())
	// func (t Time) In(loc *Location) Time 返回特定时区时间
	fmt.Println(time.Now().In(time.UTC))
	// func (t Time) Local() Time
	fmt.Println(time.Now().Local())
	// func (t Time) UTC() Time
	fmt.Println(time.Now().UTC())
	// 返回自定义时间 + 时区
	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	fmt.Println(time.Date(2018, 9, 12, 12, 30, 10, 0, time.Local))
	// func Parse(layout, value string) (Time, error) 字符串转换为时间类型
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2018-04-23 12/24/51"))

	// 字符串转换为时间格式并设置时区
	// func ParseInLocation(layout, value string, loc *Location) (Time, error) (layout已带时区时可直接用Parse)
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local))
	// Unix返回与给定Unix时间相对应的本地时间
	// func Unix(func Unix(sec int64, nsec int64) Time
	fmt.Println(time.Unix(1, 2))

}

func testTimeFormat() {
	//
	fmt.Println("Format Start")
	// func
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
	fmt.Println(time.Now().Format("2006年01月02日 15:04:05"))
	fmt.Println(time.Now().Format(time.RFC3339))
	// func (t Time) AppendFormat(b []byte, layout string) []byte
	fmt.Println(time.Now().AppendFormat([]byte("时间: "), "2006-01-02 15:04:05"))

}

func testTimeTimestamp() {
	//func (t Time) Unix() int64
	fmt.Println(time.Now().Unix())
	//func (t Time) UnixNano() int64
	fmt.Println(time.Now().UnixNano())

}

func testTimeFuncDate() {
	fmt.Println("FuncDate Start")
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().YearDay())
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().ISOWeek())
}

func testTimeOperate() {
	fmt.Println("Operate Start")
	// func (t Time) AddDate(years int, months int, days int) Time {
	fmt.Println(time.Now().AddDate(0, 0, 1))
	// func Until(t Time) Duration
	fmt.Println(time.Until(time.Now()))
	// func Since(t Time) Duration
	fmt.Println(time.Since(time.Now()))
	// func (t Time) Sub(u Time) Duration
	fmt.Println(time.Now().Sub(time.Date(2018, 9, 12, 12, 30, 10, 0, time.Local)))
	// func (t Time) Add(d Duration) Time
	fmt.Println(time.Now().Add(time.Duration(1) * time.Second))

}

func testTimeEqual() {
	fmt.Println("TimeEqual Start")
	// func (t Time) Equal(u Time) bool
	fmt.Println(time.Now().Equal(time.Now()))
	// func (t Time) Before(u Time) bool
	fmt.Println(time.Now().Before(time.Now()))
	// func (t Time) After(u Time) bool
	fmt.Println(time.Now().After(time.Now()))
}

func testTimeZone() {
	fmt.Println("Time Zone Start")
	//func (t Time) Zone() (name string, offset int) {
	fmt.Println(time.Now().Zone())
	//func (t Time) Location() *Location {
	fmt.Println(time.Now().Location())
}

func testTimeTimer() {
	fmt.Println("Timer Start")

	// func After(d Duration) <-chan Time

	// func AfterFunc(d Duration, f func()) *Timer

	// func (t *Timer) Reset(d Duration) bool

	// func NewTimer(d Duration) *Timer

	// func (t *Timer) Stop() bool
}

func testTimeTicker() {
	fmt.Println("Time Ticker")
	// func NewTicker(d Duration) *Ticker

	// func (t *Ticker) Stop()

	// func Tick(d Duration) <-chan Time
}

func testTimeOtherFunc() {
	fmt.Println("OtherFunc Start")
	// func Sleep(d Duration)
	time.Sleep(time.Duration(1) * time.Second)

	//
	encode, _ := time.Now().GobEncode()
	fmt.Println(encode)
	// func parseGMT(value string) int
}
