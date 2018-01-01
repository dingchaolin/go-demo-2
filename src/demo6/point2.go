package main

import (
	"math"
	"fmt"
)

type Point1 struct{
	X, Y float64
}

func (p *Point1)Distance( q Point1) float64{
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point1)GetX( q Point1) float64{
	return p.X
}
func DistanceTotal(path []Point1)float64{
	var dis float64
	for i := 1; i < len(path); i ++ {
		dis += path[i].Distance(path[i-1])
	}
	return dis
}


type Path []Point1

func(path Path) Distance()float64{
	var dis float64
	for i := 1; i < path.LenPoints(); i ++ {
		dis += path[i].Distance(path[i-1])
	}
	return dis
}

func (path Path) LenPoints() int{
	return len(path)
}


func main(){
	p := Point1{1,2}
	q := Point1{4,6}

	fmt.Println( p.Distance(q) )

	path := Path{Point1{1,2}, Point1{3,4}, Point1{5,6}}
	dis := DistanceTotal(path)

	fmt.Println( dis )

	dis1 := path.Distance()
	fmt.Println( dis1 )

	/*
	匿名类型
	 */
	var a struct{
		X int
		Y int
	}
	a.X = 123
}