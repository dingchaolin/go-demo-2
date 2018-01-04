package main

import "fmt"

func sum( s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main(){
	s := []int{7,2,8,-9,4,0}
	// x, y := <- c, <- c // 此处先执行会造成死锁
	c := make( chan int)
	go sum(s[:len(s)/2], c )
	go sum(s[len(s)/2: ], c )

	x, y := <- c, <- c//此处不能保证哪个先运行

	fmt.Println( x, y, x+y )
}