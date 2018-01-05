package main

import (
	"time"
	"fmt"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)// time.NewTick(1000 * time.Millisecond).C
	boom := time.After(5000 * time.Millisecond)

	for{
		select{
		case <- tick:
			fmt.Println("滴答...")
			case <- boom:
				fmt.Println("boom!!!")
				return
		default://当其他的case都不通的时候 执行default case
			fmt.Println("喝一口酒")
			time.Sleep( 100 * time.Millisecond )// 如果不是sleep 会打印很多的 "喝一口酒"
		}
	}
}
