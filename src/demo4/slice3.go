package main

import "fmt"

func main(){
	s := []int{2,3,4,5,6,7}
	printSlice( s )//len=6 cap=6  [2 3 4 5 6 7]
	fmt.Println( &s[0])

	//切的时候 只是记录起始位置 和 长度
	//只要不引起内存分配  值不会变
	s = s[:0]//len=0 cap=6  []
	printSlice( s )
	//改变了len 没有改变cap
	s = s[:4]//len=4 cap=6  [2 3 4 5]
	printSlice( s )

	//改变了len 改变了 cap
	s = s[2:]//len=2 cap=4  [4 5]
	printSlice( s )

}

func printSlice( s []int){
	fmt.Printf("len=%d cap=%d  %v\n", len(s), cap(s), s )
}