package main

import (
	"regexp"
	"fmt"
)

func main(){
	reg := regexp.MustCompile("[0-9]{2,5}")//表达式有问题编译时直接panic
	ok := reg.MatchString("123cb")
	fmt.Println( ok )

	fmt.Println( reg.FindString("acd1234"))// 获取 1234
	// go中正则表达式的时间复杂度是o(n)
}