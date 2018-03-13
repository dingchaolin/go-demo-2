package main

import (
	//"time"
	"fmt"
	"sync"
)

func main(){
	var n int32 = 0
	wg := new (sync.WaitGroup)
	wg.Add(2)
	var lock sync.Mutex
	go func (){
		lock.Lock()
		defer lock.Unlock()
		n += 2
		wg.Done()
	}()

	go func(){
		lock.Lock()
		defer lock.Unlock()
		n/=2
		wg.Done()
	}()

	wg.Wait()
	fmt.Println( n )


}
