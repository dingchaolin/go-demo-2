package main

import (
	"demo12/monitor/common"
	"fmt"
	"net"
	"encoding/json"
	"time"
	"log"
	"bufio"
)

type Sender struct {
	addr string
	ch   chan *common.Metric
}

func NewSender(addr string) *Sender {
	// 构造sender
	return &Sender{
		addr: "127.0.0.1:8080",
		ch : make(chan *common.Metric,10000),
	}
}

/*
断线重连
 */
func (s *Sender) connect() net.Conn{
	n := time.Microsecond * 100
	for{
		conn, err := net.Dial("tcp", s.addr)

		if err != nil {
			log.Println(err)
			time.Sleep(n)
			n = n * 2
			if n > time.Second * 30 {
				n = time.Second * 30
			}
			continue
		}
		return conn

	}
}

func (s *Sender) Start() {
	// 建立连接
	// 循环从s.ch里面读取metric
	// 序列化metric
	// 发送数据

	conn := s.connect()

	// 方式1
	//for{
	//	metric := <- s.ch
	//	buf, _ := json.Marshal(metric)
	//	fmt.Fprintf(conn, "%s\n", buf)
	//}
	// 方式2
	for metric := range s.ch {
		buf, _ := json.Marshal(metric)
		_, err := fmt.Fprintf(conn, "%s\n", buf)
		if err != nil{
			conn.Close()
			conn = s.connect()
		}

	}
}

/*
空间满了 会发送
时间到了 也会发送
 */
func (s *Sender) StartByTimeByCount() {

	conn := s.connect()
	log.Println(conn.LocalAddr())
	// bufio.NewWriter 是个蓄水池 默认蓄4k数据 如果没有满 就会一直蓄
	w := bufio.NewWriter(conn)//
	ticker := time.NewTicker(time.Second*5)

	for {
		select {
			case metric := <-s.ch:
				buf, _ := json.Marshal(metric)
				_, err := fmt.Fprintf(w, "%s\n", buf)
				if err != nil{
					conn.Close()
					conn = s.connect()
					w = bufio.NewWriter(conn)
					log.Println(conn.LocalAddr())

				}
			case <- ticker.C:
				//如果没有蓄满 时间到了 会强行flush
				err := w.Flush()
				if err != nil{
					conn.Close()
					conn = s.connect()
					w = bufio.NewWriter(conn)

				}

		}


	}
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}
