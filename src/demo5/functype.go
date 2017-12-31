package main

import (
	"strconv"
	"os"
	"fmt"
)

func add( m, n int ) int{
	return m + n
}

func sub(m , n int) int{
	return m - n
}

func main(){
	//var f func(int,int)int  定义一个函数类型
	funcmap := map[string]func(int,int)int{
		"+":add,
		"-":sub,
	}

	m, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi( os.Args[3])

	f, ok := funcmap[os.Args[2]]
	if ok{
		fmt.Println(f(m,n))
	}
}
