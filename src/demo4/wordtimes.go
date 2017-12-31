package main

import (
	"io/ioutil"
	"os"
	"log"
	"strings"
	"fmt"
)

func main(){
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)//打印错误 并直接退出程序
	}

	count := make( map[string]int)
	words := strings.Fields(string(content))
	for _, word := range words{
		if _, ok := count[word]; ok{
			count[word] ++
		}else{
			count[word] = 1
		}
	}

	for word, cnt := range count{
		fmt.Println( word, ":", cnt )
	}


}