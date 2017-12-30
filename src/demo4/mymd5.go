package main

import (
	"crypto/md5"
	"fmt"
)

func main(){
	var sum [16]byte //md5的表示类型
	md5sum := md5.Sum([]byte("hello"))
	fmt.Printf("%x\n", md5sum)

	md5sum1 := md5.Sum([]byte("hello1"))
	if md5sum == md5sum1{
		fmt.Println("相等")
	}

	sum = md5sum
	fmt.Printf("%x\n", sum )

}