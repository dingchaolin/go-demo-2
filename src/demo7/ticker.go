package main

import (
	"time"
	"fmt"
)

func main(){
	timer := time.NewTicker(time.Second)//1s一次的定时器
	cnt := 0

	for _ = range timer.C {
		cnt ++
		if cnt > 10{
			timer.Stop()//可以停止 不会关掉channel  只会停止发送数据
			return
		}
		fmt.Println( "hello    ", cnt )
	}
}