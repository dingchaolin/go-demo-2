package main

import "fmt"

func main(){
	var s []int
	printSlice2(s)

	s = append(s,0)//当cap不够用的时候 会开辟新的内容空间 会发生一个copy操作  返回一个新的地址
	printSlice2( s )

	s = append( s, 1)
	printSlice2(s)

	s= append(s, 2,3,4)
	printSlice2( s )

	s1 := []int{6,7,8}
	s = append( s, s1...)// s1... 就是把s1所有的元素都添加到s中去
	printSlice2( s )


}

func printSlice2( s []int){
	fmt.Printf("len=%d cap=%d  %v\n", len(s), cap(s), s )
}