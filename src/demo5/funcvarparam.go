package main

import "fmt"

func sum( args ...int) int {
	n := 0
	for i := 0; i < len(args); i ++{
		n += args[i]
	}
	return n
}

/*
...args  放前面交聚合  都聚合到args中
args...  放后面叫发散 把所有值列出来
 */
func main(){
	fmt.Println(sum(1,2,3))
	arr := []int{1,2,3,4}
	fmt.Println(sum(arr...))
}