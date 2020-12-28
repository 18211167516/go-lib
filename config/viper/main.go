package main

import (
	"fmt"
	"log"

	"reflect"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal(err)
		} else {
			log.Fatal(err)
			// Config file was found but another error was produced
		}
	}

	fmt.Println(viper.Get("log_level"), reflect.TypeOf(viper.Get("log_level")))

	fmt.Println("mysql ip: ", viper.Get("mysql.host"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql name: ", viper.Get("mysql.name"))

	fmt.Println("redis port: ", viper.Get("redis.port"), reflect.TypeOf(viper.Get("redis.port")))

	viper.Set("redis.port", "3306")

	//viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	//文件存在直接就忽略了
	//viper.SafeWriteConfig()
	//viper.WriteConfigAs("./cnf/config.toml")
	//viper.SafeWriteConfigAs("./cnf/config.toml") // will error since it has already been written
}
