package main

import (
	"path/filepath"
	"os"
	"fmt"
)
/*
遍历文件树
 */
func main(){
	filepath.Walk(".", func( path string, info os.FileInfo, err error) error{
		fmt.Println( path,  info )
		/*
		path 文件路径
		info 文件信息
		err 如果有err 可以判断 跳过 或者继续 返回error遍历就会中断
		 */

		 fmt.Println( info.IsDir() )//是不是目录 很多方法
		return nil
	})
}