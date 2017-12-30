package main

import (
	"os"
	"log"
)

func main(){

	f, err := os.Create("a.txt")//文件不存在会创建  文件存在会清空
	if err != nil {
		log.Fatal( err )
	}

	f.WriteString( "hello world ! 你好" )
	f.Close()
}