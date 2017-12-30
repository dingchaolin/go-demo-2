package main

import "fmt"

func main(){
	s1 := []int{2,3,4,5,6,7}
	//s := []int{}
	s1 = s1[1:4]
	fmt.Println( s1 )

	s1 = s1[:2]//从0开始
	fmt.Println( s1 )

	s1 = s1[1:]//到结尾
	fmt.Println( s1 )

	s1 = s1[:]//获取全部
	fmt.Println( s1 )
}
