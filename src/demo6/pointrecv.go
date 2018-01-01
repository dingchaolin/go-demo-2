package main

import "fmt"

type PO struct{
	X ,Y float64
}

func (p *PO) ScaleBy( factor float64){
	p.X *= factor
	p.Y *= factor
}


func main(){
	//直接指针
	p := &PO{1,2}
	p.ScaleBy( 2 )
	fmt.Println( p )

	//声明结构体后再用指针指向
	p1 := PO{1,2}
	p2 := &p1
	p2.ScaleBy( 3 )// 等价于 (&p1).ScaleBy( 3 )
	fmt.Println( p2 )


	//使用结构体调用 隐式取地址
	p3 := PO{1,2}
	p3.ScaleBy( 4 )
	fmt.Println( p3 )
}
