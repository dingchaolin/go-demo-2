package main

import (
	"os"
	"log"
)
// defer 是在函数结束的时候调用

func main(){
	f, err := os.Open("a.txt")
	if err != nil{
		log.Fatal(err)
	}

	/*
	一定在err判断之后再defer
	 */
	defer f.Close()//return 之前一定会执行
}