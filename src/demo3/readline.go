package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"io"
)

func main(){
	//如果每次写入都覆盖掉 使用这个参数 O_TRUNC
	//默认有都是追加写入

	//f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	f, err := os.Open("a.txt")//以纯读取的方式打开
	//O_RDWR 可读可写
	if err != nil {
		log.Fatal( err )
	}

	r := bufio.NewReader(f)
	//line,_ := r.ReadString('\n')
	//fmt.Print( line )
	//line,_ = r.ReadString('\n')
	//fmt.Print( line )

	count := 0
	for{
		line, err := r.ReadString('\n')
		if err == io.EOF {//文件结束
			break
		}
		fmt.Print(line)
		count ++
		if count >= 10{
			break
		}
	}
	f.Close()
}