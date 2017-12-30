package main

import "fmt"

func feib( index int ) int {
	if index <= 2 {
		return 1
	}else{
		return feib( index - 1 ) + feib( index - 2 )
	}
}
func main(){
	i := 1
	for{
		value := feib( i )
		if value > 100 {
			break;
		}
		fmt.Println( "index = ", i, "value====",  value)
		i ++
	}
}