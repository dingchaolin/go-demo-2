package main

import (
	"net"
	"log"
	"strconv"
	"time"
)

func handleConn( conn net.Conn, cnt int ){
	conn.Write([]byte("hello golang " + strconv.Itoa(cnt ) + " \n"))
	time.Sleep(time.Minute)//模拟服务器处理请求需要耗时
	conn.Close()//不关闭连接  会造成资源泄露
}
func main(){
	addr := ":8021"//监听任意ip的端口
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	cnt := 0
	for{
		// 接受连接
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal(err)
		}

		/*
		哪里阻塞go哪里
		 */
		go handleConn(conn, cnt)
		cnt ++
	}

}
// telnet 127.0.0.1  8021