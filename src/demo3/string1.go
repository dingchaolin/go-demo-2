package main

import "fmt"

func toupper(s string) string{
	var s1 []byte
	var c byte
	for i := 0; i < len(s); i ++{

		if s[i] >= 'a' && s[i] <= 'z' {
			c = s[i]
			c = c - ('h' - 'H')
		}

		s1 = append( s1, c )
	}
	return string(s1)

}
// \a 响值  跟电脑有关 能响
func main(){
	// 相加
	s1 := "hello " + "world"

	//取字符
	var c1 byte
	fmt.Println(0, len(s1)-1)
	c1 = s1[0]
	// s1[0] = 96 不能被修改
	fmt.Printf("%d   %c", c1, c1 ) // %d 整数 %c 表示字符  f - format

	//切片
	s2 := s1[0:3] // [start, end)  都是下表的位置 左闭右开  都可以省略  : 4 表示从0开始  4： 表示到结尾  ：表示从0到结尾


	fmt.Println( s1, c1, s2 )

	var b byte
	for b = 0 ; b < 177; b ++{
		fmt.Printf("%d  %c \n", b, b )
	}

	fmt.Println(0xa)

	//修改字符串内容
	array := []byte(s1)//[104 101 108 108 111 32 119 111 114 108 100]
	fmt.Println( array )
	array[0] = 'A'
	fmt.Println( array )
	s1 = string(array)
	fmt.Println( s1 )

	fmt.Println('a' + ('H' - 'h'))

	fmt.Println( toupper( s1 ))

}
