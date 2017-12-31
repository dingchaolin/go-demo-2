package main

import "fmt"

func main(){
	// byte => unit
	// rune => int32

	s := "golang你好" // go 中 一个汉字占用3个字节
	fmt.Println(len(s))//12  len是字节数
	cnt := 0
	for _, r := range s{//合法的utf字符数
		cnt += 1
		fmt.Printf( "%c\n", r )
	}

	fmt.Println( "cnt==", cnt )//8
	cnt = 0
	for _, r := range []byte(s){//合法的utf字符数
		cnt += 1
		fmt.Printf( "%c\n", r )
	}

	fmt.Println( "cnt==", cnt )//20

	cnt = 0
	ss := []rune("golang你好")// 专门处理utf8相关 rune->int32
	for _, r := range ss{//合法的utf字符数
		cnt += 1
		fmt.Printf( "%c\n", r )
	}
	fmt.Println( string(ss), cnt )


}