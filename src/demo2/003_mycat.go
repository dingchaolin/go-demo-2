package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	printFile(os.Args[1])
}

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
