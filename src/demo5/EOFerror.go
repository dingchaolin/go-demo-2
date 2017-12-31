package main

import (
	"os"
	"io"
	"log"
	"fmt"
)
// io.EOF
// socket 中 对方关闭socket 会得到这个error
// 文件的末尾会得到这个error
// 管道通信 关闭管道也会得到这个error
//这个是个标志 不应该说是个错误

func read1( f *os.File)( string, error){
	var total []byte

	buf := make([]byte, 1024)
	for{
		n, err := f.Read(buf)
		if err == io.EOF{
			break
		}

		if err != nil{
			return "", err
		}

		total = append(total, buf[:n]...)// n == len(buf)
		fmt.Println( total )
	}

	return string(total), nil

}


func main(){
	f, err := os.Open("a.txt")

	if err != nil{
		log.Fatal("open error:%v", err)
	}

	s, err := read1(f)
	if err != nil{
		log.Fatalf("read error:%v", err)
	}
	fmt.Println( "s=====",s )
}