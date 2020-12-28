package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	viper.Set("redis.port", "3306")

	viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	//文件存在直接就忽略了
	//viper.SafeWriteConfig()
	//viper.WriteConfigAs("./cnf/config.toml")
	//viper.SafeWriteConfigAs("./cnf/config.toml") // will error since it has already been written
}
