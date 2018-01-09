package main

import (
	"time"
	"fmt"
)

type Account struct{
	flag bool
	money int
}

func ( a *Account ) DoPrepare(){
	time.Sleep(time.Second)
}

func (a *Account) GetGongZi( n int ){
	a.money += n
}

func (a *Account) GiveWife(n int ){
	for !a.flag {
		time.Sleep( time.Millisecond )
	}
	a.flag = false
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
	a.flag = true
}

func (a *Account ) Buy( n int ){
	for !a.flag {
		time.Sleep( time.Millisecond )
	}
	a.flag = false
	if a.money > n {
		a.DoPrepare()
		a.money -= n

	}
	a.flag = true
}

func (a *Account) Left() int{
	return a.money
}

func main(){
	var account Account
	account.flag = true
	account.GetGongZi( 10 )
	go account.GiveWife( 6 )
	go account.Buy(5)
	time.Sleep( 2 * time.Second )
	fmt.Println( account.Left())
}