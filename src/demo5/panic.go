package main

import "fmt"

func print1(){
	defer func(){
		err := recover()//当前函数退出 print1退出
		fmt.Println( err )
	}()

	var p *int
	fmt.Println( *p )//空指针异常
}

func main(){
	defer func(){
		err := recover()//当前函数退出 main退出
		fmt.Println( err )
	}()
	print1()
	panic("执行不下去了")
	var n int
	fmt.Println( 10 / n )//panic: runtime error: integer divide by zero


	var slice [3]int
	//fmt.Println( slice[3])//下标异常
	fmt.Println( slice[2])
}