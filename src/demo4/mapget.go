package main

import "fmt"

func main(){
	ages := map[string]int{
		"a" : 1,
		"b" : 2,
	}
	fmt.Println(ages["a"])
	fmt.Println(ages["c"])//如果不存在 不会报错 会返回元素类型的默认值

	ages["a"] = ages["b"] + 2

	c, ok := ages["c"]
	fmt.Println(ok)// true  false
	if ok{
		fmt.Println(c)
	}else{
		fmt.Println("not found")
	}

	if c, ok := ages["c"]; ok{
		fmt.Println(c)
	}

}