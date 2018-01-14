package main

import (
	"net"
	"log"
	//"strconv"
	//"time"
)
var content = `HTTP/1.1 200 OK
Date: Sat, 13 Jan 2018 07:36:40 GMT
Content-Type: text/html
Connection: Keep-Alive
Server: BWS/1.1
X-UA-Compatible: IE=Edge,chrome=1
BDPAGETYPE: 3
Set-Cookie: BDSVRTM=0; path=/

<html>
<body >
<h1 style="color:red">hello golang</h1>
</body>
</html>
`
func handleConn( conn net.Conn, cnt int ){
	conn.Write([]byte(content))
	//time.Sleep(time.Minute)//模拟服务器处理请求需要耗时
	/*
	此种类型都是短连接  连接上之后就断掉了
	 */
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