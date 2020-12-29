package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("toml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
	log_level = "re"
	bool = true
	[mysql]
	  host = "yourhost"
	  name = "yourname"
	  password = "yourpassword"
	  user = "youruser"
	
	[redis]
	  port = 3306
	
	[time]
	  a = "2020-12-13 02:20:03"
	  d = "40ms"
	
`)

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	fmt.Println(viper.Get("mysql.password")) // this would be "steve"
}
