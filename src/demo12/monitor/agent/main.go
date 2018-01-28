package main

import (
	"demo12/monitor/common"
	"os"
	"github.com/shirou/gopsutil/cpu"
	"time"
	"runtime"
	"flag"
	"github.com/shirou/gopsutil/load"
)

var (
	transAddr = flag.String("trans", "127.0.0.1:80", "trans address")
)

func NewMetric( metric string, value float64) *common.Metric{
	hostname, _ := os.Hostname()
	return &common.Metric{
		Metric:    metric,
		EndPoint:  hostname,
		Value:     value,
		Tag:       []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}
}
func getData(metricStr string)*common.Metric{

	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	metric := NewMetric(metricStr, cpus[0])
	return metric

}

func CpuMetric() []*common.Metric{
	var ret []*common.Metric
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	metric := NewMetric("cpu.useage", cpus[0])
	ret = append(ret, metric)

	cpuload , err := load.Avg()
	if err == nil{
		metric = NewMetric("cpu.load1", cpuload.Load1)
		ret = append(ret, metric)
		metric = NewMetric("cpu.load5", cpuload.Load5)
		ret = append(ret, metric)
	}
	return ret
}

func main() {

    flag.Parse()
	addr := *transAddr
	sender := NewSender(addr)
	//go sender.Start()
	ch := sender.Channel()
/*
	ticker := time.NewTicker(time.Second * 5)
	metric := getData("cpu.usage")
	for range ticker.C {
		ch <- metric
	}
*/
	sched := NewSched(ch)
	sched.AddMetric( CpuMetric, time.Second)
	sender.Start()

}
