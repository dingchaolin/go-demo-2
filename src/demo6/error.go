package main

import (
	"errors"
	"fmt"
)

func main(){
	var e error
	e = errors.New( "an error")
	var cmd string
	e = fmt.Errorf("bad command:%s", cmd)
	fmt.Println( e )
	fmt.Println( e.Error() )

	/*
	EOF
	var EOF= errors.New("EOF")
	 */


}