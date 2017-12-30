package main

import "fmt"

func main(){
	a := make([]int, 5)
	printSlice1("a", a )

	b := make([]int, 0, 5)
	printSlice1( "b", b )

	c := b[0:2]
	printSlice1( "c", c )

	d := c[2:5]//cap跟实际能访问的元素个数有有关 cap >= len
	printSlice1( "d", d )//d len=3 cap=3  [0 0 0]

}

func printSlice1( s string, x []int){
	fmt.Printf("%s len=%d cap=%d  %v\n",s, len(x), cap(x), x )
}