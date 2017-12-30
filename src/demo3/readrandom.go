package main

import (
	"os"
	"log"
	"fmt"
)

func main(){
	//如果每次写入都覆盖掉 使用这个参数 O_TRUNC
	//默认有都是追加写入

	//f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	f, err := os.Open("a.txt")//以纯读取的方式打开
	if err != nil {
		log.Fatal( err )
	}
	f.Seek(3, os.SEEK_SET )
	buf := make([]byte, 2)
	f.Read(buf)
	fmt.Println( string(buf))
	f.Close()
}