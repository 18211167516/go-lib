package main

import (
	"fmt"
	"time"
)

//h 小时
func Example_duartion_Hour() {
	tp, _ := time.ParseDuration("11h")
	fmt.Println(tp.Hours())
	// Output: 11
}

//分钟
func Example_duartion_Minutes() {
	tp, _ := time.ParseDuration("10m")
	fmt.Println(tp.Minutes())
	// Output: 10
}

//s 秒
func Example_duartion_Seconds() {
	tp, _ := time.ParseDuration("10s")
	fmt.Println(tp.Seconds())
	// Output: 10
}

// ms 毫秒
func Example_duartion_Milliseconds() {
	tp, _ := time.ParseDuration("10ms")
	fmt.Println(tp.Milliseconds())
	// Output: 10
}

// us 微秒
func Example_duartion_Microseconds() {
	tp, _ := time.ParseDuration("1us")

	fmt.Println(tp.Microseconds())

	// Output: 1
}

//ns 纳秒
func Example_duartion_Nanoseconds() {
	tp, _ := time.ParseDuration("100ns")
	fmt.Println(tp.Nanoseconds())
	// Output: 100
}

/*

  func (d Duration) Round(m Duration) Duration {

  }

  返回距离d最近的时间段，得到的是晚于d的时间，
  该时间段应该满足能整除m；
  如果有两个满足要求的时间点，距离d相同，会向上舍入；
  如果m <= 0，会返回d的拷贝。


  M 参数示例
  time.Nanosecond,  //按纳秒四舍五入，9位  0.0000000000 小数点后保留9位
  time.Microsecond, //按微秒，6位 0.000000 小数点后保留6位
  time.Millisecond, //按毫秒，3位，0省略 0.000 小数点后保留3位
  time.Second,      //按秒
  2 * time.Second,  //按2秒
  time.Minute,      //按分
  10 * time.Minute, //按10分
  time.Hour,
*/
func Example_duartion_Round_Nanosecond() {
	tp, _ := time.ParseDuration("10508888788.9ns")
	fmt.Println(tp.Round(time.Second))
	// Output:11s
}

func Example_duartion_Truncate() {
	tp, _ := time.ParseDuration("10508888788.9ns")
	fmt.Println(tp.Truncate(time.Second))
	// Output:10s
}

func Example_duartion_second() {
	second := time.Duration(10) * time.Second
	fmt.Println(second)
	// Output:10s
}

func Example_duartion_hour() {
	hour := time.Duration(2) * time.Hour
	fmt.Println(hour.Hours())
	// Output:2
}
