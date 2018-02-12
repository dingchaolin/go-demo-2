package main

import (
	"demo13/common"
	"os"
	"github.com/shirou/gopsutil/cpu"
	"time"
	"runtime"
	"flag"
	"github.com/shirou/gopsutil/load"
	"os/exec"
	"strconv"
	"bufio"
	"strings"
	"log"
	"demo13/agent"
	"fmt"
)

var (
	transAddr = flag.String("trans", "127.0.0.1:6001", "trans address")
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

func NewUserMetric(cmdstr string) agent.MetricFunc{
	return func()[]*common.Metric {
		metrics, err := getUserMetrics(cmdstr)
		if err != nil{
			log.Print(err)
			return []*common.Metric{}
		}
		return metrics
	}
}
func getUserMetrics( cmdstr string)([]*common.Metric, error){
	//构建命令
	//获取标准输出
	//按行解析
	//获取key, value
	//包装成common.Metric
	var ret []*common.Metric
	cmd := exec.Command( "bash", "-c", cmdstr)
	stdout, _ := cmd.StdoutPipe()

	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(stdout)
	for{
		line, err := r.ReadString('\n')
		if err != nil{
			break
		}
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)
		if len( fields ) != 2 {
			continue
		}
		key, value := fields[0], fields[1]
		n, err := strconv.ParseFloat( value, 64)
		if err != nil{
			log.Print(err)
			continue
		}
		metric := NewMetric( key, n )
		ret = append(ret, metric)
	}
	return ret , err

}
func main() {

    flag.Parse()
	addr := *transAddr
	fmt.Println( addr )
	sender := agent.NewSender(addr)
	//go sender.Start()
	ch := sender.Channel()
/*
	ticker := time.NewTicker(time.Second * 5)
	metric := getData("cpu.usage")
	for range ticker.C {
		ch <- metric
	}
*/
	sched := agent.NewSched(ch)
	sched.AddMetric( CpuMetric, time.Second)
	sched.AddMetric(NewUserMetric("./usr.py"), 3*time.Second)
	sender.Start()

}
