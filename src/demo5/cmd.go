package main

import (
	"os/exec"
	"log"
	"fmt"
)

func main(){
	cmd := exec.Command("ls", "-l")
	out, err := cmd.CombinedOutput()//返回标准输出和标准错误一起 不用wait 内部已经做了处理
	if err != nil {
		log.Fatal( err )
	}

	fmt.Println( string(out) )

}