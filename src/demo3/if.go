package main

import (
	"strconv"
	"fmt"
)

func main(){
	s := "abc123"

	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println( err )
		return
	}else{
		fmt.Println( n )
	}

	if n, err := strconv.Atoi( s ); err != nil {
		fmt.Println( n )
	}
}