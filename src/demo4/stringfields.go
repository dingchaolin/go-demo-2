package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println( strings.Fields(" bbb ccc     dddd"))//返回一个数组 按空格分割 [bbb ccc dddd]
}