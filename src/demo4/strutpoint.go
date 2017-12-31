package main

import (
	"fmt"
	"unsafe"
)

type Student1 struct{
	Id   int
	Name string
}

func main(){
	s1 := Student1{
		Id : 2,
		Name : "dcl",
	}

	var p *Student1
	p = &s1
	p.Id = 444

	fmt.Println( s1 )

	var p1 *int
	p1 = &s1.Id
	*p1 = 555

	fmt.Println( s1 )

	fmt.Println( unsafe.Sizeof(s1.Name))// 16

	//var arr [3]Student //数组类型结构体
	//var ss []Student //slice类型结构体
	//var m map[string]Student //map类型结构体
}