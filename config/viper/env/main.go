package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	// 绑定环境变量
	viper.AutomaticEnv()
}

func main() {
	viper.BindEnv("go.path", "GOPATH", "GOOS")
	fmt.Println("GOPATH: ", viper.Get("GOPATH"))
	fmt.Println("GOPATH: ", viper.Get("go.path"))

	//设置env前缀
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")

	os.Setenv("SPF_ID", "13") // typically done outside of the app

	fmt.Println(viper.Get("id")) // 13

	viper.BindEnv("empty")
	os.Setenv("SPF_empty", "")
	fmt.Println(viper.AllSettings()) // 13
	viper.AllowEmptyEnv(true)
	fmt.Println(viper.AllSettings()) // 13

	os.Setenv("SPF_a_rep", "222")
	replacer := strings.NewReplacer("-", "_")

	viper.SetEnvKeyReplacer(replacer)
	fmt.Println(viper.Get("a-rep")) // 13
}
