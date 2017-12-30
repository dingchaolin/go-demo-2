package main

import (
	"os"
	"log"
	"fmt"
)

func main(){

	f, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)//n表示最大读取多少个 -1 表示所有
	for _,info := range infos {
		fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name() )
	}
    fmt.Println("==================================================")
	names, err := f.Readdirnames( -1 )
	for _,name := range names {
		fmt.Printf("%s\n", name )
	}
	f.Close()
}