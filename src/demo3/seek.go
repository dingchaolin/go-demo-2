package main

import (
	"os"
	"log"
)

func main(){
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
	//O_RDWR 可读可写
	if err != nil {
		log.Fatal( err )
	}

	f.WriteString( "hello world ! 你好" )
	f.Seek(10, os.SEEK_SET)//开始位置， 跳过2个字符 写入
	f.WriteString( "$" )
	f.Close()
}