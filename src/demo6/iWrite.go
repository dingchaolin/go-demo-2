package main

import (
	"io"
	"os"
	"fmt"
)

type ByteCounter int

func (b *ByteCounter) Write( p []byte)(int, error){
	*b += ByteCounter(len(p))
	return int(*b), nil
}

func main(){
	b := new( ByteCounter)
	io.Copy( b, os.Stdin )
	fmt.Println( *b )
}

//go run iWrite.go < iWrite.go