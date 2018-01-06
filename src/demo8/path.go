package main

import (
	"fmt"
	"path"//处理ftp http相关的
	//"path/filepath"//处理不同操作系统相关的
	"path/filepath"
)

func main() {
	s := "/aa/bb/cc/README.md"
	fmt.Println(path.Dir(s))//获取该文件的目录
	fmt.Println(path.Base(s))//获取该文件的文件名
	fmt.Println(path.Ext(s))//获取该文件的拓展名
	dir := path.Dir(s)
	name := path.Base(s)
	fullname := filepath.Join( dir, name )
	fmt.Println( fullname )

}
