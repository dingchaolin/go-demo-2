package main

import (
	"net"
	"log"
	//"strconv"
	//"time"
	"os"
	"io/ioutil"
	"bufio"
	"strings"
	"flag"
)

var (
	root = flag.String("root", "/", "root of ftp server data dir")
)
// client -> GET /home/dingchaolin/a.txt\n
// server -> content of /home/dingchaolin/a.txt

// client -> STORE /home/dingchaolin/a.txt\n content of file
// server -> OK

// client -> LS /home/dingchaolin/a.txt\n content of file
// server -> content of dir /home/dingchaolin
func handleConnFtp( conn net.Conn ){
	// 读取客户端需要的文件
	// 从conn中读取一行内容
	// 按空格分割指令和文件名

	// 打开文件
	// 读取内容
	// 发送内容
	// 关闭连接和文件
	defer conn.Close()//不关闭连接  会造成资源泄露

	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2 {
		conn.Write([]byte("bad input"))
		return
	}
	cmd := fields[0]
	name := fields[1]

	if cmd == "GET"{
		f, err := os.Open(name )
		if err != nil{
			log.Println( err )
			return
		}
		defer f.Close()
		buf, err := ioutil.ReadAll(f)
		if err != nil{
			log.Println(err)
			return
		}
		conn.Write([]byte(buf))
	}else if cmd == "STORE"{
		// 从 r 中读取文件内容
		// 创建name文件
		// 向文件写入数据
		// 往conn写入ok
		// 关闭连接和文件
	}


}
func main(){
	addr := ":8021"//监听任意ip的端口
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for{
		// 接受连接
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal(err)
		}

		/*
		哪里阻塞go哪里
		 */
		go handleConnFtp(conn)
	}

}
// telnet 127.0.0.1  8021