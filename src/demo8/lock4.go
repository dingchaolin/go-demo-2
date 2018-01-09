package main

import (
	"time"
	"fmt"
	"sync"
)

/*
互斥锁
 */
type A struct{
	flag sync.Mutex//初始化是未锁住的状态
	money int
}

func ( a *A ) DoPrepare(){
	time.Sleep(time.Second)
}

func (a *A) GetGongZi( n int ){
	a.money += n
}

func (a *A) GiveWife(n int ){
	a.flag.Lock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
	defer a.flag.Unlock()
	//a.flag.Unlock()
}

func (a *A ) Buy( n int ){
	a.flag.Lock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n

	}
	//a.flag.Unlock()
	defer a.flag.Unlock()
}

func (a *A) Left() int{
	return a.money
}

func main(){
	var account A
	account.GetGongZi( 10 )
	wg := new (sync.WaitGroup)
	wg.Add(2)
	go func(){
		account.GiveWife( 6 )
		wg.Done()
	}()
	go func() {
		account.Buy(5)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println( account.Left())
}