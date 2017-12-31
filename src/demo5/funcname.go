package main

import (
	"fmt"
	"os"
)

func split( sum int)(x, y int){//不赋值就是默认值 int的默认值是0
	x = sum/10
	y = sum%10
	return
}

func main(){
	fmt.Println(split(12))
	os.Exit( 3 )//退出进程
}
