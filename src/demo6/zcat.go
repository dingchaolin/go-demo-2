package main

import (
	"compress/gzip"
	"os"
	"fmt"
	"io"
)

/*
过滤器对象 - 中间件
 */
func main(){
	r,err := gzip.NewReader( os.Stdin )
	if err != nil {
		fmt.Println( err )
		return
	}
	io.Copy( os.Stdout, r )


}

// zcat < XX.gzip