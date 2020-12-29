package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.Int("mysql.host", 8381, "Mysql host")
	pflag.Parse()
	// 绑定命令行
	viper.BindPFlags(pflag.CommandLine)

	fmt.Println(viper.GetInt("mysql.host"))
}
