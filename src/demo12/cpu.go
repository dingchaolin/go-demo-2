package main

import (
	"time"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

func main(){
	//计算cpu的使用率
	cpus, err := cpu.Percent(time.Second, true )// 是否计算每一个cpu的使用率  true 所有 false某一个
	if err != nil{
		panic( err )
	}
	fmt.Println(cpus)
    // load 能反映cpu是否繁忙  负载  1 表示满负荷运行 高于1 cpu任务已经排队了
	loadavg, err := load.Avg()
	if err != nil{
		panic( err )
	}
	fmt.Println( loadavg )

	memstat, err := mem.VirtualMemory()
	if err != nil{
		panic(err)
	}
	fmt.Println( memstat.UsedPercent )
	fmt.Println( memstat.Used )// 单位: 位


	diskstat, err := disk.Usage("/")
	if err != nil{
		panic( err )
	}
	fmt.Println( diskstat.UsedPercent )

}