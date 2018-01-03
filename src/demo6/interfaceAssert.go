package main

import (
	"io"
	"os"
	"bytes"
)

func main(){
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)//类型断言的语法 w.(原来的类型) 相当于通过这一步把原来的类型找回来
	c,ok:= w.(*bytes.Buffer)
	_ = f
	_ = c
/*
	var v int
	ret , ok := v.(int)
	if ok {

	}
	ret , ok = v.( uint )
*/
	/*
	直到把所有的类型都试一下 找到真正的类型
	 */
}