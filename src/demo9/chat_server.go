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

var globalRppm *Room = NewRoom()

type Room struct{
	users map[string]net.Conn
}

func NewRoom() *Room{
	return &Room{
		users:make(map[string]net.Conn),
	}
}
func (r *Room) Join( user string, conn net.Conn){
	conn, ok := r.users[user]
	if ok {
		r.Leave(user)
	}
	r.users[user] = conn
}

func (r *Room)Leave( user string){
	//关掉连接
	// 从users里面删除
}

func (r *Room)Broadcast(user string, msg string){
	// 遍历所以的用户 发送消息

}
// client -> binggan 123456
// client -> hello golang
// client -> close

//接受新的连接
//验证用户的账号和密码
//等待用户的输入
//向所有在线的用户广播用户的输入

func chatHandleConn(conn net.Conn){
	defer conn.Close()
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len( fields ) != 2{
		conn.Write([]byte("bad input"))
		return
	}

	user := fields[0]
	password := fields[1]
	if password != "123456"{
		return
	}
	// join用户
	for{
		//获取用户输入
		//broadcast
	}
	//leave用户


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