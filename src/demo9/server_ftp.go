package main

import (
	"net"
	"log"
	//"strconv"
	//"time"
	"os"
	//"io/ioutil"
	"bufio"
	"strings"
	"flag"
	"io"
	"path/filepath"
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

	r := bufio.NewReader(conn)//会进行超前读取
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
		//大文件
		// 方法1
		//buf := make([]byte, 4096)
		//for{
		//	n, err := f.Read(buf)
		//	if err != io.EOF{
		//		break
		//	}
		//	conn.Write(buf[:n])
		//}

		// 方法2
		// 一行能代替以上的所有代码
		//块读取
		// 把文件中的数据写到socket中
		io.Copy(conn, f)
		/*
		socket中的数据写到文件中
		 */
		//io.Copy(conn, f)

		//小文件
		//buf, err := ioutil.ReadAll(f)
		//if err != nil{
		//	log.Println(err)
		//	return
		//}
		//conn.Write([]byte(buf))
	}else if cmd == "STORE"{
		// 从 r 中读取文件内容
		// 创建name文件
		// 向文件写入数据
		// 往conn写入ok
		// 关闭连接和文件
		os.MkdirAll(filepath.Dir(name), 0755)
		f, err := os.Create(name)
		if err != nil{
			log.Println( err )
			return
		}
		io.Copy(f, r )// 是r 不是conn
		f.Close()
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