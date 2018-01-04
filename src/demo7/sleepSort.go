package main

import (
	"time"
	"fmt"
)
/*
sleep 排序
 */
func main(){
	s := []int{2,7,1,6,4}
	for _, n := range s{
		go func( n int ){
			time.Sleep( time.Duration(n) * time.Second )
			fmt.Println( n )
		}(n)
	}
	time.Sleep( 10 * time.Second )
}
