package main

import (
	"os"
	"log"
)

func main(){
	//如果每次写入都覆盖掉 使用这个参数 O_TRUNC
	//默认有都是追加写入

	f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	//O_RDWR 可读可写
	if err != nil {
		log.Fatal( err )
	}

	f.WriteString( "hello world ! 你好" )
	f.Close()
}