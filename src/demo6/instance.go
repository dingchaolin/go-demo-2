package main

import (
	"fmt"
	"math"
)

type IDistance interface{
	Distance() float64
}

/*
对象中实现了接口中的所有方法 就算是实现了这个接口
 */

 type POINT struct {
 	X float64
 	Y float64
 }

 func (point POINT) Distance() float64{
 	return math.Hypot( point.X, point.Y )
 }

func (point POINT) Distance2Point( q POINT ) float64{
	return math.Hypot( q.X - point.X, q.Y -  point.Y )
}

 type PATH []POINT
 func (p PATH) Distance() float64 {
 	var sum float64
 	for i := 0; i < len(p) - 1; i++ {
 		sum += p[i].Distance2Point(p[i+1])
	 }
 	return sum
 }

func main(){

	path := make(PATH, 3 )
	p1 := POINT{1,2}
	p2 := POINT{3,4}
	p3 := POINT{5,6}

	path[0] = p1
	path[1] = p2
	path[2] = p3

	var i IDistance
	i = p1
	fmt.Println( i.Distance() )

	i = p2
	fmt.Println( i.Distance() )

	i = path
	fmt.Println( path.Distance() )


}