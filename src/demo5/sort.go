package main

import (
	"sort"
	"fmt"
)

type Student struct{
	Name string
	Id int
}
/*
排序 只能针对slice的 数组不行
 */
func main(){
	s := []int{2, 3, 1, 5, 9, 7}

	sort.Slice( s, func( i, j int) bool{
		return s[i] < s[j]//正序排列
		//return s[i] > s[j]//倒序排列
	})

	fmt.Println( s )

	ss := []Student{}
	ss = append(ss, Student{
		Name: "aa",
		Id: 2,
	})

	ss = append(ss, Student{
		Name: "bb",
		Id: 3,
	})

	ss = append(ss, Student{
		Name: "cc",
		Id: 1,
	})

	sort.Slice( ss, func( i, j int)bool{
		return ss[i].Id < ss[j].Id
	})

	fmt.Println( ss )

	sArr := []string{ "dd", "bb", "aa", "cc"}

	sort.Strings( sArr )

	fmt.Println( sArr )


}