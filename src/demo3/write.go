package main

import (
	"os"
	"log"
	"fmt"
)
/*
文件格式化写入
 */
func main(){
	f, err := os.Create("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint( f, "hello")//写入到文件中
	fmt.Fprintln(f,"hello ln" )
	s := "hello"
	n := 4
	fmt.Fprintf( f, "my string is: %s, n = %d \n", s, n )

	f.Close()
}