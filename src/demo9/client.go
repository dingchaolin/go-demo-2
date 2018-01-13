package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"os"
)
// tcp 是双工的 任何一方都可以主动发起关闭  关闭之后 另一方会收到 EOF
func main(){
	//addr := "www.baidu.com:80"
	addr := "127.0.0.1:8021"
	//拨号
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( conn.RemoteAddr().String() )
	fmt.Println( conn.LocalAddr().String() )

	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("write size  ", n )

	//buf := make([]byte, 4096)
	//buf := make([]byte, 10)
	//n, err = conn.Read(buf)//读写都是流式数据
	///*
	//EOF 在网络编程中 表示对方关闭了链接 唯一判断对方关闭了链接
	// */
	//if err != nil && err != io.EOF{
	//	log.Fatal(err)
	//}
	//fmt.Println(string(buf[:n]))
	io.Copy( os.Stdout, conn)

	conn.Close()
}

// dig baidu.com 查看所有的dns ip

