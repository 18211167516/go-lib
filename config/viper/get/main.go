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
		log.Fatal(err)
	}

	fmt.Println("bool: ", viper.GetBool("bool"))
	fmt.Println("redis port int: ", viper.GetInt("redis.port"), reflect.TypeOf(viper.GetInt("redis.port")))
	fmt.Println("redis port float: ", viper.GetFloat64("redis.port"), reflect.TypeOf(viper.GetFloat64("redis.port")))
	fmt.Println("redis port GetIntSlice: ", viper.GetIntSlice("redis.port"), reflect.TypeOf(viper.GetIntSlice("redis.port")))
	fmt.Println("redis port GetString: ", viper.GetString("redis.port"), reflect.TypeOf(viper.GetString("redis.port")))
	fmt.Println("mysql GetStringMap", viper.GetStringMap("mysql"), reflect.TypeOf(viper.GetStringMap("mysql")))
	fmt.Println("mysql GetStringMapString", viper.GetStringMapString("mysql"), reflect.TypeOf(viper.GetStringMapString("mysql")))
	fmt.Println("mysql GetStringSlice", viper.GetStringSlice("mysql"), reflect.TypeOf(viper.GetStringSlice("mysql")))

	//
	fmt.Println("time GetTime", viper.GetTime("time.a"), reflect.TypeOf(viper.GetTime("time.a")))
	fmt.Println("time GetDuration", viper.GetDuration("time.d"), reflect.TypeOf(viper.GetDuration("time.d")))

	fmt.Println("key is exits", viper.IsSet("key.a"))
	fmt.Println("allkey", viper.AllSettings())

}
