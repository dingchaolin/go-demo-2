package main

import (
	"net"
	"log"
	"bufio"
	"errors"
	"io"
	"encoding/binary"
	"fmt"
)

func mustReadByte( r *bufio.Reader) byte{
	b, err := r.ReadByte()
	if err != nil{
		panic( err )
	}
	return b
}
// 1. 握手  客户端发给服务器一串东西 服务器再返回一串东西
// 2. 获取客户端的代理请求
// 3. 开始代理
func readAddr(r *bufio.Reader)(string, error){
	defer func(){
		e := recover()// 返回 interface{}类型 断言成功 就能得到真实的类型 失败会继续panic 只有接口能够断言
		err := e.(error)//  类型断言
		fmt.Println( err )

	}()
	version := mustReadByte(r)//版本
	if version != 5 {
		return  "",errors.New("bad version")
	}
	log.Printf("version:%d", version )
	cmd, _ := r.ReadByte()
	if cmd != 1 {// 建立连接
		return  "",errors.New("bad cmd")
	}
	log.Printf("cmd:%d", cmd )

	mustReadByte(r)//保留字段跳过 rsv

	addrtype,_ := r.ReadByte()
	log.Printf("addrtype:%d", addrtype )// 0 表示不需要验证 2 表示用户名密码验证
	if addrtype != 3{//域名
		return  "",errors.New("bad addrtype")
	}

	// 读取一个字节的数据 代表厚点紧跟着域名的长度
	// 读取n个字节得到域名 n根据上一步得到的结果来决定
	addrlen, _ := r.ReadByte() //地址的长度
	addr := make( []byte, addrlen)
	io.ReadFull( r, addr )
	log.Printf( "addr===%s", string(addr) )

	var port int16
	binary.Read(r, binary.BigEndian, &port )//网络上都是大端序
	log.Printf( "port===%d", port )


	return fmt.Sprintf("%s:%d", addr, port), nil
}
func handshake( r *bufio.Reader, conn net.Conn ) error {

	version, _ := r.ReadByte()//版本
	log.Printf("version:%d", version )
	if version != 5 {
		return  errors.New("bad version")
	}
	nmethods, _ := r.ReadByte()//客户端支持几种验证方式
	log.Printf("nmethods:%d", nmethods )

	buf := make([]byte, nmethods)// method 的长度就是nmethods的内容
	io.ReadFull(r, buf )//将空间填充满
	log.Printf("methods:%d", buf )// 0 表示不需要验证 2 表示用户名密码验证

	resp := []byte{5, 0 }//版本号 认证方式
	conn.Write( resp )
	return nil
}
func HandleConn( conn net.Conn ){
	defer conn.Close()
	r := bufio.NewReader( conn )
	handshake(r, conn )
	addr, _ := readAddr( r )
	fmt.Println("======addr====", addr )// https port 443
	// resp
	conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
}
func main(){
	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal( err )
	}
	for{
		 // accept new connection
		 conn, _ := l.Accept()
		 go HandleConn(conn)
	}
}
