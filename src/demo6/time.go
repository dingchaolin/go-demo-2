package main

import (
	"fmt"
	"time"
)

func main(){
	var n time.Duration
	n = time.Millisecond * 2
	fmt.Println( n.String())
	fmt.Println( n.Seconds() )
	time.Sleep( n )
	fmt.Println( time.Now() )

	t := time.Now()
	t1 := t.Add(-time.Hour)//一小时之前
	fmt.Println( t )
	fmt.Println( t1.Sub( t ))

	fmt.Println( time.Since( t ))


}
