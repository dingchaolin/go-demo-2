package main

import "fmt"

/*
select 多路选择 哪个可用 就走哪个
 */
func fibonacci2( c, quit chan int ){
	x, y := 0, 1
	for{
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}

	}
}

/*
先接收 后发送
 */
func main(){
	c := make( chan int )
	quit := make(chan int)
	/*
	协程退出的条件：函数退出 协程就退出了
	 */
	go func(){
		for i:= 0; i < 30; i ++ {
			fmt.Println( <- c)
		}
		quit <- 1
	}()

	fibonacci2( c, quit )
}