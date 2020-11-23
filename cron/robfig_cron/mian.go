package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

type testJob struct{}

func (t testJob) Run() {
	panic("test job")
	//fmt.Println("i.m test job")
}

func main() {
	
	f, _ := os.Create("cron.log")

	logger := cron.VerbosePrintfLogger(log.New(io.MultiWriter(f, os.Stdout), "cron: ", log.LstdFlags))
	c := cron.New(cron.WithChain(cron.Recover(logger)), cron.WithSeconds(), cron.WithLogger(logger))
	c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(testJob{}))
	c.AddFunc("*/2 * * * * *", func() {
		fmt.Println("test 11")
	})

	spce := fmt.Sprint("@every ", time.Duration(1)*time.Second)
	c.AddFunc(spce, func() {
		//panic("1232132")
	})

	//c.AddJob("* * * * * *", testJob{})

	c.Start()

	time.AfterFunc(time.Second*5, func() {
		c.Stop()
	})
	//time.Sleep(time.Second * 12)
	select {}
}
