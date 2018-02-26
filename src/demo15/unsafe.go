package main

import (
	"fmt"
	"unsafe"
)

type T struct{
	m int8
	n int64
}

func main(){
	var n int
	fmt.Println( unsafe.Sizeof(n)) //变量占用的字节数 8
	var t T
	fmt.Println( unsafe.Sizeof(t)) //变量占用的字节数 16 对齐
	fmt.Println( unsafe.Alignof(t)) // 8 对齐的字节数
	fmt.Println( unsafe.Offsetof(t.n)) // n在t中的偏移量8

	var m [2]int8

	n ++

	var p *int = &n

	*p = 32

	//p = &m[0] //类型不配置 会报错
	// 任何指针可以转化为unsafe指针 unsafe指针可以转为为任何指针
	p = (*int)(unsafe.Pointer(&m[0]))//这样可以解决上面的报错问题
	*p = 0x1010
	fmt.Println(m)

}
