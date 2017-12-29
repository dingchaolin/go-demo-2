package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep1 = flag.String("s", " ", "separator")

var sep2 = flag.Bool("n", false, "换行")


func main(){

	flag.Parse()//这一行不要忘记 解析命令行

	fmt.Println( strings.Join(flag.Args(), *sep1))

	if *sep2  {
		fmt.Println()
	}
}