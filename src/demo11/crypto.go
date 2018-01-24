package main

import (
	"io"
	"net"
	"reflect"
	"fmt"
)

type CryptoWriter struct{
	w io.Writer
}

func NewCryptoWriter(w io.Writer, key string) io.Writer{
	return &CryptoWriter{
		w: w,
	}
}

func ( w *CryptoWriter) Write(b []byte)(int, error){

}

type CryptoReader struct{
	r io.Reader,
}

func NewCryptoReader(r io.Reader, key string) io.Reader{
	return &CryptoReader{
		r:r,
	}
}


func ( r *CryptoReader) Read(b []byte)(int, error){

}
func main(){
	key := "123456"
	remote, err := net.Dial("tcp", "")
	if err != nil{
		fmt.Println( err )
		return 
	}
	w := NewCryptoWriter( remote, key)
	w.Write([]byte("hello"))

	r := NewCryptoReader(remote, key)
	buf := make([]byte, 1024)
	r.Read(buf)
}