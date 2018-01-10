package main

import (
	"net/http"
	"log"
	"sync"
)

func work( ch chan string, wg *sync.WaitGroup){
	for u := range ch{
		//u := <- ch
		resp, err := http.Get(u)
		if err != nil {
			log.Print( err )
			return
		}
		log.Printf( "%s : %d", u,resp.StatusCode)
		resp.Body.Close()
	}
	wg.Done()



}

func main(){
	wg := new(sync.WaitGroup)
	wg.Add(5)
	taskch := make(chan string)
	for i := 0; i < 5; i ++ {
		go work(taskch, wg)
	}

	urls := [] string{"https://www.baidu.com", "http://news.baidu.com/"}
	for _, url := range urls{
		taskch <- url
	}
	close(taskch)
	wg.Wait()
}