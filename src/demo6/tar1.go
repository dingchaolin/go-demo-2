package main

import (
	"archive/tar"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"compress/gzip"
)

func main(){
	//uncompress,err := gzip.NewReader( os.Stdin )
	//tr := tar.NewReader( uncompress )//解压gzip
	//上面可以实现解压gzip

	tr := tar.NewReader( os.Stdin )//解压tar


	hdr, err := tr.Next()

	if err != nil {
		return
	}

	fmt.Println( hdr.Name )
	//info := hdr.FileInfo()
	io.Copy( ioutil.Discard, tr )
}