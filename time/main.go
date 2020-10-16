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
	testTimeZone()
	testTimeEqual()
	testTimeOperate()
	testTimeTimer()
	testTimeOtherFunc()
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
	fmt.Println(string(time.Now().AppendFormat([]byte("时间: "), "2006-01-02 15:04:05")))
}

func testTimeTimestamp() {
	fmt.Println("Timestamp Start")
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
	// func (t Time) AddDate(years int, months int, days int) Time
	fmt.Println(time.Now().AddDate(0, 0, 1))
	// func (t Time) Add(d Duration) Time 当前时间加传入时间段
	fmt.Println(time.Now().Add(time.Duration(1) * time.Hour))
	// func Until(t Time) Duration 传入时间减本地时间
	fmt.Println(time.Until(time.Date(2018, 9, 12, 12, 30, 10, 0, time.Local)))
	// func Since(t Time) Duration 本地时间减传入时间
	fmt.Println(time.Since(time.Date(2018, 9, 12, 12, 30, 10, 0, time.Local)))
	// func (t Time) Sub(u Time) Duration 当前时间减传入时间
	fmt.Println(time.Now().Sub(time.Date(2018, 9, 12, 12, 30, 10, 0, time.Local)))

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
	//func (t Time) Zone() (name string, offset int) 算在时间t生效的时区，返回缩写
	fmt.Println(time.Now().Zone())
	//func (t Time) Location() *Location 获取当前时间设置的时区
	fmt.Println(time.Now().Location())
	//func (t Time) In(loc *Location) Time 设置时区返回时间
	fmt.Println(time.Now().In(time.UTC))
	fmt.Println(time.Now().In(time.UTC).Zone())

	// func (t Time) Local() Time
	fmt.Println(time.Now().Local())
	// func (t Time) UTC() Time
	fmt.Println(time.Now().UTC())

	// func LoadLocation(name string) (*Location, error) 加载时区
	location, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC).In(location))
	// func FixedZone(name string, offset int) *Location 设置时区及其偏移量
	loc := time.FixedZone("UTC-8", -8*60*60)
	fmt.Println(time.Date(2018, 9, 12, 12, 30, 10, 0, loc))
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", loc))
}

func testTimeTimer() {
	fmt.Println("Timer Start")
	/* timer := time.NewTimer(time.Second * 2)

	//多路复用通道
	select {
	case <-timer.C: //计时器到时了，即2秒已到
		fmt.Println("time is over,stop!!")
	case <-time.After(time.Second * 1): //打点器触发了，说明已隔500毫秒

		fmt.Println("已超时")
	}
	*/
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

func testTimeDuration() {

	tp, _ := time.ParseDuration("1.5s")
	fmt.Println(tp.Truncate(1000), tp.Seconds(), tp.Nanoseconds())

	//func (d Duration) Hours() float64
	//func (d Duration) Minutes() float64
	//func (d Duration) Seconds() float64
	//func (d Duration) Nanoseconds() int64
	//func (d Duration) Round(m Duration) Duration         // 四舍五入
	//func (d Duration) Truncate(m Duration) Duration      // 向下取整
}

func testTimeOtherFunc() {
	fmt.Println("OtherFunc Start")
	//func (t Time) IsZero() bool
	fmt.Println(time.Now().IsZero())
	// func Sleep(d Duration)
	time.Sleep(time.Duration(1) * time.Second)
	// func (t Time) GobEncode() ([]byte, error)
	encode, _ := time.Now().GobEncode()
	fmt.Println(encode)
	// func (t *Time) GobDecode(data []byte) error
	var ignored time.Time
	fmt.Println(ignored.GobDecode(encode))
	fmt.Println(string(encode))
	//marshal binary序列化，将时间t序列化后存入[]byte数组中
	//func (t Time) MarshalBinary() ([]byte, error)

	//marshal json序列化，将时间t序列化后存入[]byte数组中
	//func (t Time) MarshalJSON() ([]byte, error)

	//marshal text序列化，将时间t序列化后存入[]byte数组中
	//func (t Time) MarshalText() ([]byte, error)

	//将data数据反序列化到时间ｔ中
	//func (t *Time) UnmarshalBinary(data []byte) error

	//将data数据反序列化到时间ｔ中
	//func (t *Time) UnmarshalJSON(data []byte) (err error)

	//将data数据反序列化到时间ｔ中
	//func (t *Time) UnmarshalText(data []byte) (err error)

	tbyte, err := time.Now().MarshalJSON()
	fmt.Println("'t.MarshalJSON': ", tbyte, err)

	// 时间数据反序列化
	var tun time.Time
	err = tun.UnmarshalJSON(tbyte)
	fmt.Println("'t_un.UnmarshalJSON': ", tun, err)
}
