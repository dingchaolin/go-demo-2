package main

import (
	"fmt"
	"time"
	//"os"
)

func test(ch chan int , id int ){
	exitCode := 0
	go ctrl( ch, id, &exitCode)
	go consume(id, &exitCode)

}

func consume( id int, exit *int){

	for{
		if *exit == 1{
			fmt.Println("协程  ", id , "退出")
			return
		}
		time.Sleep( time.Second * 1)
		fmt.Println("协程~~~~~~~", id )
	}
}

func ctrl(ch chan int , id int, code *int ){
	for{
		ret := <-ch
		fmt.Println( id, "收到！", ret)
		if ret == 1 {
			*code = 1
			//os.Exit( 0 )//main会退出
			return  //当前协程退出
			//close( ch )

		}
	}

}

func main(){
	ch := map[int] chan int{}
	ch[0] = make(chan int)
	ch[1] = make(chan int)

	go test(ch[0], 1)
	go test(ch[1], 2)

	time.Sleep( 5 * time.Second )

	//ch <- 1
	ch[0] <- 1
	ch[1] <- 2

	time.Sleep( 10 * time.Second )


}