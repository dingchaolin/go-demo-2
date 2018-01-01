package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"os/exec"
)

func main(){
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("chaolindeMBP:~ chaolinding@%s$", host)

	r := bufio.NewScanner(os.Stdin)//按行读取 也 可以按词读取
	//r := bufio.NewReader(os.Stdin)
	for{
		fmt.Print(prompt)//打印前置文本
		if !r.Scan(){
			break
		}

		line := r.Text()

		//line , _ := r.ReadString('/n')
		//line = strings.TrimSpace( line )
		if len(line) == 0 {
			continue
		}


		args := strings.Fields(line)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			fmt.Println(err)
		}
	}
}