package main

import (
	"os"
	"log"
	"bufio"
	"io/ioutil"
	"fmt"
)

func main(){
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	//裸读取 很少使用
	buf := make([]byte, 4096)
	n ,err := f.Read(buf)
	buf = buf[:n]

	// 加上buffer的读取 和高效
	r := bufio.NewReader( f )
	r.Read(buf)

	// 按行读取 按分隔符读取
	r1 := bufio.NewScanner( f )
	line := r1.Text()
	fmt.Println( line )

	//小文件
	ioutil.ReadFile("a.txt")//不打开文件
	ioutil.ReadAll(f)//要打开文件

	//类文件类型  内存 管道 socket...
	//操作文件的神器
	//copy.go

}