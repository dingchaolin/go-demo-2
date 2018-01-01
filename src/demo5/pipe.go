package main

import (
	//"io"
	"os/exec"
	"os"
	"log"
	"strings"
)

func main(){

	line := "cat pipe.go|grep main"
	cmds := strings.Split(line, "|")
	s1 := strings.Fields(cmds[0])
	s2 := strings.Fields(cmds[1])

	cmd1 := exec.Command(s1[0], s1[1:]...)
	cmd1.Stdin = os.Stdin
	out, _ := cmd1.StdoutPipe()//cmd1 结束的时候 会关闭管道
	/*
	管道中的数据如果没人消费 数据会积压
	 */
	cmd2 := exec.Command(s2[0], s2[1:]...)
	cmd2.Stdin = out
	cmd2.Stdout = os.Stdout
	cmd1.Start()// run 是阻塞的
	cmd2.Start()

	log.Print("start" )
	cmd1.Wait()
	cmd2.Wait()
}