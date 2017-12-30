package main

import (
	"math/rand"
	"fmt"
)

func main(){
	for i:=0; i < 10; i ++{
		fmt.Println( rand.Intn( 10 )) // 0 - 9
	}
}