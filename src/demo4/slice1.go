package main

import "fmt"

func main(){
	names := [4]string{
		"a",
		"b",
		"c",
		"d",//go 里面强制要求最后一个加,
	}

	fmt.Println( names )

	a := names[0:2]
	b := names[1:3]
	fmt.Println( a, b )

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println( names)
	fmt.Println("------------------------")
	// c切的时候 不用管a的引用是谁， 只管a有什么  a只有2个元素 下标是 0， 1 当c切a的1，2的时候 只能切到1 所有 c之后一个元素
	// 二次切的时候 以当前切片为基准
	// 修改的值都是修改的引用数组的值
	c := a[1:2]
	c[0] = "YYYY"

	fmt.Println("a====", a, "\n b======", b,"\n c=======", c)
	fmt.Println( names)
}