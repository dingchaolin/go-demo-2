package main

import (
	"fmt"
	"errors"
)

func iter( s []int) func() (int, error){
	var i = 0
	return func() (int, error){
		if i > len(s){
			return 0, errors.New("end")
		}
		n := s[i]
		i ++
		return n, nil
	}
}


func main(){
	f := iter([]int{1,2,3})
	for{
		n, err := f()
		if err != nil{
			fmt.Println( "err====", err )
			break
		}else{
			fmt.Println(n)
		}

	}

}