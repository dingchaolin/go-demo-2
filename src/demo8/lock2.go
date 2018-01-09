package main

import (
	"time"
	"fmt"
	"sync"
)

/*
互斥锁
 */
type Acc struct{
	flag sync.Mutex//初始化是未锁住的状态
	money int
}

func ( a *Acc ) DoPrepare(){
	time.Sleep(time.Second)
}

func (a *Acc) GetGongZi( n int ){
	a.money += n
}

func (a *Acc) GiveWife(n int ){
	a.flag.Lock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
	defer a.flag.Unlock()
	//a.flag.Unlock()
}

func (a *Acc ) Buy( n int ){
	a.flag.Lock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n

	}
	//a.flag.Unlock()
	defer a.flag.Unlock()
}

func (a *Acc) Left() int{
	return a.money
}

func main(){
	var account Acc
	account.GetGongZi( 10 )
	go account.GiveWife( 6 )
	go account.Buy(5)
	time.Sleep( 2 * time.Second )
	fmt.Println( account.Left())
}