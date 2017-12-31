package main

import (
	"strings"
	"fmt"
)

func toupper( s string) string{
	return strings.Map( func( r rune) rune{
		return r - ('a' - 'A')
	}, s)
}

func main(){
	fmt.Println( toupper("hello"))
}