package main

import (
	"fmt"
	"unsafe"
)

//slice 由3部分组成  struct
// data pointer
// len int64
// cap int64
// reflect 中有定义 reflect.SliceHeader
type SliceHeader struct{
	Data unsafe.Pointer
	Len int64
	Cap int64
}

// 自定义slice
func slice(s []int, begin int64, len int64)[]int{
	hdr := SliceHeader{
		Data: unsafe.Pointer(&s[begin]),//unsafe.Pointer(&s),//有错
		Len: len,
		Cap: int64(cap(s)),
	}

	s1 := *(*[]int)(unsafe.Pointer(&hdr))
	return s1
}

type StringHeader struct{
	Data unsafe.Pointer
	Len int
}


func stringSlice( s string, begin int, len int)string{
	// 进行指针的加减运算之前 要转成uintptr
	//hdr := StringHeader{
	//	Data: unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(begin)),
	//	Len: len,
	//}
	// 不能修改原来string的数据
	hdr := *(*StringHeader)(unsafe.Pointer(&s))
	hdr.Data = unsafe.Pointer(uintptr(unsafe.Pointer(hdr.Data)) + uintptr(begin))
	hdr.Len = len
	s1 := *(*string)(unsafe.Pointer(&hdr))
	return s1
}

func ZeroCopyString( buf []byte )string{
	hdr := &StringHeader{
		Data: unsafe.Pointer(&buf[0]),
		Len: len(buf),
	}
	return *(*string)(unsafe.Pointer(hdr))
}
func main(){
	s := []int{1,2,3}
	fmt.Println(&s[0])

	s1 := s[0:1]//切片本质就是改变了len cap 其他都没变
	var p *SliceHeader
	p = (*SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("%#v\n", *p)
	p = (*SliceHeader)(unsafe.Pointer(&s1))
	fmt.Println("%#v\n", *p)

	hdr := SliceHeader{
		Data: unsafe.Pointer(&s[0]),//unsafe.Pointer(&s),//有错
		Len: 2,
		Cap: 3,
	}

	s3 := *(*[]int)(unsafe.Pointer(&hdr))
	fmt.Println( len(s3) )
	fmt.Println( cap(s3) )
	fmt.Println( s3[0], s3[1] )

	fmt.Println( slice(s, 1, 2))

	str := "hellonihao"
	str1 := stringSlice(str, 0, 2 )
	fmt.Println( str1 )


	buf := []byte{'h','e', 'l', 'l', 'o'}
	str2 := string(buf)
	buf[0] = 'a'
	fmt.Println( string(buf), str2)// 没有共享内存 发生了拷贝

	buf1 := []byte{'h','e', 'l', 'l', 'o'}
	str3 := ZeroCopyString(buf)
	buf1[0] = 'a'
	fmt.Println( string(buf1), str3)// 共享内存 没有发生拷贝

}