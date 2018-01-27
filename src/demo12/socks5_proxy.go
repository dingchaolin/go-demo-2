package main

import (
	"net"
	"log"
	"bufio"
	"errors"
	"io"
	"encoding/binary"
	"fmt"
	"sync"
	"crypto/rc4"
	"crypto/md5"
)

/*
面向接口编程
最小接口原则
 */
type CryptoWriter struct{
	w io.Writer
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer{
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil{
		panic(err)
	}
	return &CryptoWriter{
		w: w,
		cipher:cipher,
	}
}
// 把b里面的数据进行加密 之后写入到w.w中
// 调用w.w.Write进行写入
func ( w *CryptoWriter) Write(b []byte)(int, error){
	buf := make([]byte, len(b))
	w.cipher.XORKeyStream(buf,b)
	return w.w.Write(buf)
}

type CryptoReader struct{
	r io.Reader
	cipher *rc4.Cipher
}

func NewCryptoReader(r io.Reader, key string) io.Reader{
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil{
		panic(err)
	}
	return &CryptoReader{
		r:r,
		cipher:cipher,
	}
}


func ( r *CryptoReader) Read(b []byte)(int, error){
	n, err := r.r.Read(b)
	buf := b[:n]
	r.cipher.XORKeyStream(buf,buf)
	return n, err

}

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
func handshake( r *bufio.Reader, w io.Writer ) error {

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
	w.Write( resp )
	return nil
}
func HandleConn( conn net.Conn ){
	defer conn.Close()
	r := bufio.NewReader( NewCryptoReader(conn, "123456") )
	w := NewCryptoWriter( conn, "123456")


	handshake(r, conn )
	addr, _ := readAddr( r )
	fmt.Println("======addr====", addr )// https port 443
	// resp
	w.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	//代理
	// 建立到目标服务器的连接
	var remote net.Conn
	var err error
	remote, err = net.Dial("tcp",  addr)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	wg := new (sync.WaitGroup)
	wg.Add(2)
	// go 读取客户端的数据 发送到remote 直到conn的EOF 关闭remote
	// src 源关闭了 io.Copy(dest， src) 才会返回
	go func(){
		defer wg.Done()
		io.Copy(remote, r )
		remote.Close()
	}()
	// go 读取remote的数据 发送到客户端conn 直到remote的EOF 关闭conn
	go func(){
		defer wg.Done()
		io.Copy(w, remote)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()
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
