package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var (
		x0 int //跟系统的长度一致 受编辑器和操作系统影响 在网络传输上 尽量不要使用这个类型 要使用一个精确的类型
		x1 int32
		x2 int64
		x3 uint//跟系统的长度一致
		x4 uint32
		x5 uint64
		x6 byte
		x7 uint8
		x8 int8
	)
	//x7 = 256 //溢出
	x8 = 127
	x8 += 1 //-128 编译器未发现溢出
	x7 = 255 //溢出
	x7 += 1  // 0  +=78 = 77
	fmt.Println( "x0 int length===", unsafe.Sizeof( x0 ), x0 )//8 0
	fmt.Println( "x1 int32 length===", unsafe.Sizeof( x1 ), x1 )//4 0
	fmt.Println( "x2 int64 length===", unsafe.Sizeof( x2 ), x2 )//8 0
	fmt.Println( "x3 uint length===", unsafe.Sizeof( x3 ), x3 )// 8 0
	fmt.Println( "x4 uint32 length===", unsafe.Sizeof( x4 ), x4 )// 4 0
	fmt.Println( "x5 uint64 length===", unsafe.Sizeof( x5 ), x5 )// 8 0
	fmt.Println( "x6 byte length===", unsafe.Sizeof( x6 ), x6 )// 1 0
	fmt.Println( "x7 byte length===", unsafe.Sizeof( x7 ), x7 )// 1 0
	fmt.Println( "x7 byte length===", unsafe.Sizeof( x8 ), x8 )// 1 0



}
