package main

import (
	"os"
	"io"
	"log"
	"fmt"
)

func v1(){
	var f *os.File
	var err error
	if len(os.Args) > 1 {

		f, err = os.Open(os.Args[1])
		fmt.Println( "os.Args[1]===",os.Args[1],"err====",err)
		if err != nil {
			log.Fatal( err )
			defer f.Close()
		}
	}else{
		f = os.Stdin
	}
	buf := make([]byte, 1024)
	for{
		n ,err := f.Read(buf)
		fmt.Println( "read=====",n, string(buf[:n]) )
		if err != io.EOF{
			return
		}
		os.Stdout.Write( buf[:n])
		}
}

func v2(){
	var f *os.File
	var err error
	if len(os.Args) > 1 {
		f, err = os.Open(os.Args[1])
		fmt.Println( "os.Args[1]===", os.Args[1], "err====",err)
		if err != nil {
			log.Fatal( err )
			defer f.Close()
		}
	}else{
		f = os.Stdin
	}
	io.Copy( os.Stdout, f )
}
func main(){
	v1()
	//v2()
}