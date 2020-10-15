package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds(), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))

	c.AddFunc("*/2 * * * * *", func() {
		file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE, 0755)
		defer file.Close()
		fmt.Println("test 11")
		file.Write([]byte("test 111\r\n"))
	})

	c.AddFunc("*/5 * * * * *", func() {
		time.Sleep(time.Second * 6)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("test 222")
	})

	c.Start()

	time.AfterFunc(time.Second*5, func() {
		c.Stop()
	})
	//time.Sleep(time.Second * 12)
	select {}
}
