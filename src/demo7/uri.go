package main

import (
	"net/url"
	"log"
	"fmt"
)

func main(){
	urlStr := `http://dcl:dcl@blog.csdn.net/baalhuo/article/details/51178154?name=dcl/#1`
	//urlStr := ftp://user:password@xxx.com/
	u, err := url.Parse( urlStr )
	if err != nil {
		log.Fatal( err )
	}

	fmt.Println( "schema===", u.Scheme )
	fmt.Println( "host=====", u.Host )
	fmt.Println( "path=====", u.Path )
	fmt.Println( "queryStrting====", u.RawQuery )
	fmt.Println( "user&password====", u.User )
	fmt.Println( "锚点=====", u.Fragment )
}