package main

import (
	//"fmt"
	"os"
	//"log"
	"os/exec"
)

func main(){
	cmd := exec.Command("ls", "-l")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
}