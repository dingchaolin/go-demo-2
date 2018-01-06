package main

import (
	"net/http"
	"log"
	"io"
	"os"
)

func main() {
	// http://www.baidu.com/index/name=dcl/#1
	// schema :// host / path / queryString # 锚点
	url := "http://www.baidu.com"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal( err )
	}

	defer resp.Body.Close()//此处需要关闭 否则会造成资源泄露
	// 2XX 正常  3XX 重定向 4XX 客户端错误 5XX 服务器错误
	if resp.StatusCode != http.StatusOK {
		log.Fatal( resp.Status )
	}
	io.Copy(os.Stdout, resp.Body )

}
