package main

import (
	//"fmt"
	"os"
	"log"
	"os/exec"
)

func main(){
	//fmt.Println( os.Stdin )//只读的方式打开
	//fmt.Println( os.Stdout )//只写的方法打开
	//fmt.Println( os.Stderr )//只写的方式打开
	//
	//os.Stdout.WriteString("hello world")
	//
	//os.Stderr.WriteString( "std err")

	//f, err := os.OpenFile("/dev/null", os.O_WRONLY, 0755)//输出的结果都不要了
	f ,err := os.Create("ls.out")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
    //将标准输出重定向到文件中
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = f
	cmd.Start()
	cmd.Wait()
}