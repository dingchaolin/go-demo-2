package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "separator")
/*
s 表示命令行 命令  -s
第二个参数 是 默认值
第三个参数是 --help 时候使用的
./004_myechoflag --help
Usage of ./004_myechoflag:
  -s string
        separator (default " ")

 */

func main(){
	flag.Parse()//这一行不要忘记 解析命令行
	fmt.Println( strings.Join(flag.Args(), *sep))
}