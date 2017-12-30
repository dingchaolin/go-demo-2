package main

import "fmt"

func main(){
	var b bool
	b = true
	b = false
	b = ("aa" == "bb")

	if b {
		fmt.Println( "真的")
	}else{
		fmt.Println( "假的")
	}

	fmt.Println( 3/2 ) // 1
	fmt.Println( 3/2.0 ) // 1.5
}