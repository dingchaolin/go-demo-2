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

	//删除元素
	delete( ages, "a")
	fmt.Println(ages["a"])

	//空map
	var m1 map[string]int
	//未进行初始化的map是可以跟nil进行比较的
	//一定要进行make创建之后才能使用
	if m1 == nil {
		m1 = make( map[string]int)
	}
}