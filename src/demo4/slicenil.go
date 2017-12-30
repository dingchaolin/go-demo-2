package main

import "fmt"

func main(){
	a := [...]int{2,3,4,5,6,7}
	var s []int //空切片用nil表达
	//不允许进行复制操作
	fmt.Println( s, len(s), cap(s) )
	if s == nil {
		fmt.Println( "nil" )
	}
	//赋值过一次之后 就不是nil了
	s = a[:0]
	fmt.Println(s == nil )
	fmt.Println( s, len(s), cap(s) )//[] 0 6 a切片的属性

}
