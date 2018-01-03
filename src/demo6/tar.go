package main

import (
	"io"
	"archive/tar"
	"path/filepath"
	"os"
)

func untar( base string, r io.Reader) error{
	tr := tar.NewReader( r )//读取tar文件
	/*
	会有一个文件头
	 */
	for{
		hdr, err := tr.Next()//读取文件头header
		if err == io.EOF{
			return nil
		}
		if err != nil{
			return err
		}

		fullpath := filepath.Join(base, hdr.Name)//文件的全路径
		info := hdr.FileInfo()

		//as dir
		if info.IsDir() {//创建父目录
			os.MkdirAll( fullpath, 0755 )
			continue
		}

		dir := filepath.Dir( fullpath )
		os.MkdirAll(fullpath, 0755)

		//as file
		f, err := os.Create( fullpath )
		if err != nil {
			f.Close()
			return err
		}

		_,err = io.Copy(f, tr )

		if err != nil{
			f.Close()
			return err
		}

		f.Chmod( info.Mode() )//设置属性
		f.Close()

	}
}

func main(){
	untar(".", os.Stdin )
}