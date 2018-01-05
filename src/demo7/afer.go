package main

import (
	"time"
	"fmt"
)

func main(){
	c := time.After( time.Second * 3 )//3秒之后传数据
	<- c
	fmt.Println( "done ")
}