package main

/*
空接口可以接受任何类型
 */

type Writer interface {

 	Write( b []byte)(int,error)

 }

 type I interface{

 }
func main(){

	var i I
	var n int
	i = n
	var s string
	i = s
	_ = i

}