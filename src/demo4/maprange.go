package main

import "fmt"

func main(){
	ages := map[string]int{
		"a" : 1,
		"b" : 2,
	}

	//key-value都遍历
	//遍历顺序不确定
	for name, age := range ages{
		fmt.Println("name===", name, "age===", age)
	}

	//key遍历
	for name := range ages {
		fmt.Println( name )
	}
}
