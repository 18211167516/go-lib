package main

import (
	"fmt"
	"time"
)

func ExampleNewTicker() {
	/* tick := time.Tick(1 * time.Second)
	for next := range tick {
		fmt.Printf("%v\n", next)
	} */
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}

}

func ExampleNewTimer() {

	fmt.Println("秒：", time.Now().Second())
	timer := time.NewTimer(time.Second * 5)
	time.AfterFunc(time.Second, func() {
		//该匿名函数的作用就是在1秒后打印结果，并通知main()函数可以结束主goroutine
		fmt.Println("one second after,秒：", time.Now().Second())
		timer.Reset(time.Second * 2)
	})

	//多路复用通道
	select {
	case <-timer.C: //计时器到时了，即2秒已到
		fmt.Println(timer.Stop())
		fmt.Println("time is over,stop!!秒：", time.Now().Second())
	case <-time.After(time.Second * 2): //打点器触发了，说明已隔500毫秒
		fmt.Println(timer.Stop())
		fmt.Println("已超时；秒：", time.Now().Second())
	}

	// OutPut: 已超时2
}
