package main

import (
	"os"
	"io/ioutil"
	"log"
	"time"
	"fmt"
)

func read( f *os.File)(string, error){
	buf, err := ioutil.ReadAll(f)
	if err != nil{
		return "", err

	}

	return string(buf), nil
}

func main(){
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	var content string
	retries := 3

	for i:= 1; i <= retries; i ++{
		content, err = read(f)
		if err == nil{
			break
		}
		time.Sleep( time.Second  )
	}
	fmt.Println( content )
}

/*
### 1. 出错退出进程 - 初始化的时候
- 比如初始化进程数据的时候，如果失败直接打印错误并退出

### 2. 重试的时候， 比如http请求
- 如果err != nil 就continue 接着重试 直至重试次数用完

### 3. 函数中，把err返回/上抛
- 函数中不处理err ，把err上抛， 让调用者处理

 */