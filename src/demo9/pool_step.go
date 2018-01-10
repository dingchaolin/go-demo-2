package main

import (
	"net/http"
	"log"
	"fmt"
)

// 给定一个url返回的url的status
// www.baidu.com 200 OK
func printUrl( url string){
	resp, err := http.Get( url )
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println( url, resp.Status)
}

func main(){
	printUrl("http://www.baidu.com")
}