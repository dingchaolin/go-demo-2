package main

import (
	"os/exec"
	"io"
	"fmt"
	"bufio"
	"log"
)

func main(){
	cmd := exec.Command("ls", "-l")

	out, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil{
		log.Fatal("start error: %v", err)
	}

	f := bufio.NewReader(out)

	for{
		line,err := f.ReadString('\n')
		if err == io.EOF{
			break
		}

		if err != nil{
			fmt.Println("read error: %v", err)
			break
		}
		fmt.Println( line )

	}

	err := cmd.Wait()//可以获取退出码 退出信息 处理僵尸进程
	if err != nil {
		fmt.Println( err )
	}


}