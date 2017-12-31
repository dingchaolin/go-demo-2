package main

import "fmt"
/*
通过参数传入的 就是值传递 当时是什么值 就是什么值
通过变量传入 就是引用
 */
func main(){
	var flist [] func()

	for i := 0; i < 3 ; i ++{
		flist = append( flist, func(){
			fmt.Println( i , &i)//3次打印 同一地址
		})

	}

	for i := 4; i < 7 ; i ++{
		i := i
		flist = append( flist, func(){
			fmt.Println( i , &i)//3次打印 同一地址
		})

	}

	for _, f := range flist{
		f()
	}
}