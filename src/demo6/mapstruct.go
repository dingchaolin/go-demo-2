
package main

import "fmt"

type Student struct{
	Id int
	Name string
}

func main(){
	m := make(map[string]*Student)
	/*
	struct 类型的 存指针才能修改  普通类型 int 不用指针方式也能修改
	结构体比较大的时候 指针更好
	 */
	m["dcl"] = &Student{
		Id: 2,
		Name: "dingchaolin",
	}

	m["dcl"].Id = 4

	fmt.Println( *m["dcl"] )


}
