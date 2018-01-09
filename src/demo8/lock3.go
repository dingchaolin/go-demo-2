package main

import (
	"time"
	"fmt"
	"sync"
)

/*
互斥锁
 */
type Ac struct{
	flag sync.Mutex//初始化是未锁住的状态
	money int
}

func ( a *Ac ) DoPrepare(){
	time.Sleep(time.Second)
}

func (a *Ac) GetGongZi( n int ){
	a.money += n
}

func (a *Ac) GiveWife(n int ){
	a.flag.Lock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
	defer a.flag.Unlock()
	//a.flag.Unlock()
}

func (a *Ac ) Buy( n int ){
	a.flag.Lock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n

	}
	//a.flag.Unlock()
	defer a.flag.Unlock()
}

func (a *Ac) Left() int{
	return a.money
}

func main(){
	var account Ac
	account.GetGongZi( 10 )
	ch := make(chan int)
	go func(){
		account.GiveWife( 6 )
		ch <- 0
	}()
	go func() {
		account.Buy(5)
		ch <- 0
	}()
	/*
	超时控制 100毫秒退出
	 */
	deadline := time.After(time.Millisecond * 100)
	for i := 0; i < 2; i ++{
		select{
			case <- ch:
				case <- deadline:
					fmt.Println("deadline")
					return //退出
		}

	}
	fmt.Println( account.Left())
}