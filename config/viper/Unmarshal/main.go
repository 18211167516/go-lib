package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Host    string
	Name    string
	PathMap string `mapstructure:"path_map"`
}

var C config

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	err := viper.UnmarshalKey("mysql", &C)
	//err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatal("unable to decode into struct, %v", err)
	}

	fmt.Println(C)
}
