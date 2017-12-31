package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	//var m, n int
	//fmt.Scanf("%d %d", &m, &n) // Scanln  读取一行
	//fmt.Println( m + n )

	f := bufio.NewReader( os.Stdin )
	line,err := f.ReadString('\n')//获取一整行
	fmt.Println( line, err )
}