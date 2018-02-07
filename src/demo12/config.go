package main

import (
	"github.com/BurntSushi/toml"
	"flag"
	"log"
)
type config struct{
	Sender SenderConfig
	UserScript []UserScriptConfig `toml:"user_script"`//配置文件中的字段名可以与此对应
}

type UserScriptConfig struct{
	Path string `toml:"path"`
	Step int `toml:"step"`
}

type SenderConfig struct{
	TransAddr string `toml:"trans_addr"`
	FlushInterval int `toml:"flush_interval"`
	MaxSleepTime int  `toml:"max_sleep_time"`
}
var (
	configPath = flag.String("config","./config.toml", "config path")
	gcfg config
)

/*
[] 用于结构体
[[]]用于数组
 */
func main(){
	flag.Parse()
	_, err := toml.DecodeFile(*configPath, &gcfg)
	if err != nil{
		log.Fatal(err)
	}

	log.Printf( "%#v", gcfg)
}