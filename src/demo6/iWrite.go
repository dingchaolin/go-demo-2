package main

import (
	"io"
	//"os"
	"fmt"
	"bytes"
)

type ByteCounter int

func (b *ByteCounter) Write( p []byte)(int, error){
	*b += ByteCounter(len(p))
	return int(*b), nil
}

type LineCounter struct{
	Sum int
}

func ( l *LineCounter ) Write(p []byte)( int, error){
	for _, b := range p{
		if b == '\n'{
			l.Sum++
		}
	}
	return len(p), nil
}




func main(){
	b := new( ByteCounter)
	l := new( LineCounter )

	buf := new (bytes.Buffer)

	w := io.MultiWriter(l, b)
	buf.WriteString(`
	fsa
	fdsa
	fdsa
	fqtew
	gsafds`)

	//io.Copy( w, os.Stdin )//文件读取
	io.Copy( w, buf )//内存读取
	fmt.Println( *b )
	fmt.Println( l.Sum )
}

//go run iWrite.go < iWrite.go