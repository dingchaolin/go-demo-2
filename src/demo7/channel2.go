package main

import (
	"fmt"
	"time"
)

/*
同一协程中的channel通信 必须使用带缓冲的信道
两个协程进行通信，可以使用非缓冲信道 发送和 接收是异步的
 */
func main() {
	ch := make(chan int,3)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep( 5 * time.Second )
}
