package client

import (
	"net"
	"flag"
	"io"
	"log"
	"sync"
)

var (
	target = flag.String("target", "www.baidu.com:80", "target host")
)
func handleConn( conn net.Conn ){
	// 建立到目标服务器的连接
	var remote net.Conn
	var err error
	remote, err = net.Dial("tcp", *target )
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
/*
defer conn.Close()
defer remote.Close()

go io.Copy(remote,conn)
io.Copy(conn, remote)
return
 */
	wg := new (sync.WaitGroup)
	wg.Add(2)
	// go 读取客户端的数据 发送到remote 直到conn的EOF 关闭remote
	// src 源关闭了 io.Copy(dest， src) 才会返回
	go func(){
		defer wg.Done()
		io.Copy(remote, conn )
		remote.Close()
	}()
	// go 读取remote的数据 发送到客户端conn 直到remote的EOF 关闭conn
	go func(){
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()
}
func main(){
	flag.Parse()
	//建立 listen
	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal( err )
	}
	for{
		 // accept new connection
		 conn, _ := l.Accept()
		 go handleConn(conn)
	}
}
// go run proxy.go --target=www.qq.com:80
// curl -v 127.0.0.1:8021
// go run proxy.go --target=www.qq.com:80
// ssh 22  telnet 23

// go run proxy.go --target=towel.blinkenlights.nl:23
// telnet 127.0.0.1 8021