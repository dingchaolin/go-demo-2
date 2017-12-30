package main

import (
	"fmt"
	"os"
)

func main(){

	for i := 0; i < 3; i ++{
		fmt.Println( i )
	}

	//相当于while
	i := 5
	for i < 7 {
		fmt.Println( i )
		i ++
	}

	for j, arg := range os.Args {//从1 开始打印
		fmt.Println(j, arg )
	}

	for _, arg := range os.Args {//从1 开始打印
		fmt.Println( arg )
	}

	for j, arg := range "dfghjk" {
		fmt.Println(j, arg )
	}

	//死循环
	for{
		i ++
		if( i > 10 ){
			break
		}
	}

	s := "hello中文"//一个中文3个字节
	for i, arg := range s[2:]{
		fmt.Printf("%d, %c\n",i,  arg )
	}
}