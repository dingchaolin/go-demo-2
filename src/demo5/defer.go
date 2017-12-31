package main

import (
	"fmt"
	"log"
)

func print(){
	defer func(){
		fmt.Println("defer")
	}()

	log.Fatal("出错了")//有这句的时候 defer执行不了

	fmt.Println("hello")
}

func main(){
	print()
}