package main

import (
	"net/http"
	"log"
	"fmt"
	"sync"
)

/*
适合短小的批量任务
 */
// 给定一个url返回的url的status
// www.baidu.com 200 OK
func printUrl( url string){
	resp, err := http.Get( url )
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println( url, resp.Status)
}

func _work( ch chan string, wg * sync.WaitGroup ){
	defer wg.Done()//done会让计数器减一
	/*
		for{
			url,ok  := <- ch
			if !ok {//能取到 ok为true 不能取到 ok为false 没有数据了 说明channel关闭了
			//如果channel被关闭掉了 url会得到一个channel的传输类型的默认值  并且永远不会阻塞 会取到默认值
				break
			}
			printUrl( ur)
		}
	*/
	//跟上面的等价
	for url := range ch{//当channel关闭的时候 循环会退出
		//for range 等价于上面的死循环
		printUrl( url )
	}

}
//channel特性
// 1. 只要不close 可以永远发送和接收数据
// 2. 如果channel里面没有数据 接收方会阻塞
// 3. 如果没有人正在等待channel的数据 发送方会阻塞
// 4. 从一个close的channel取数据永远不会则是 同时获取的是默认值
// 5. 一个channel中的数据是可以被多个协程竞争获取的
// 主协程启动一个work协程 同时传递一个channel
// 主协程向channel里面发送一个url
// work协程从channel里面获取url 之后调用printUrl打印url
// work一直工作直到channel关闭
// 永远不要使用sleep的方式来进行协程同步
//启动多个work协程
// 主协程向channel里面发送多个url 之后调用printUrl打印url
//work携程不停重复第三条 指导channel关闭

//创建一个WaitGroup
//调用add
//调用Wait等待携程结束

func main(){
 ch := make(chan string)//channel是一个动态数组 数组长度随时都在变化
 wg := new (sync.WaitGroup )
 // add 的次数 跟 调用done的次数要一样
 // waitgroup 只管次数 不管对应关系
 wg.Add( 3 ) //先add为好 add会让计数器家加数
 for i := 0; i < 3; i ++{
	 go _work(ch, wg )
 }
 // 起一个协程 add(1) 跟上面的效果相同
/*
	for i := 0; i < 3; i ++{
		wg.Add( 1 )
		go _work(ch, wg )
	}
*/
 for i:= 0; i < 5; i ++ {
	 url := "http://www.baidu.com"
	 ch  <- url
 }
 //在生产者中关闭通道
 close(ch)
 wg.Wait()
 //time.Sleep( time.Second * 3)
}