package main

import "fmt"

// \a 响值  跟电脑有关 能响
func main(){
	str1 := "hello  \\ \" \n \a"
	doc := `
	即使换行也不影响
	可以验证一下
	类似Nodejs的模板字符串
	`
	fmt.Println( str1, doc )

}
