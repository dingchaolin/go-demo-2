package main

import (
	"fmt"
	"unsafe"
)

func main(){
	a1 := [3]int{1,2,3}
	fmt.Println( a1 )

	var a2 [3]int
    a2 = a1//是值拷贝 不是引用
    fmt.Println(a2)

    fmt.Println( a1 == a2, &a1[0], &a2[0], unsafe.Sizeof(a1) )// 地址不同
    // 相等的前提 是 所有的值都相等
    fmt.Printf("%x\n", 255)//16进制

}