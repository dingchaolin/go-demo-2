package main

import (
	"io"
	"os"
	"bytes"
	"fmt"
)

func main(){
	var w io.Writer
	w = os.Stdout
	w = new( bytes.Buffer )
	w = nil
	if w == nil {
		fmt.Println( "nil" )
	}
}