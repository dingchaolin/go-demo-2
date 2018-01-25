package main

import (
	"testing"
	"bytes"
	"fmt"
	"strings"
	"io/ioutil"
)

/*
单元测试
 */
func TestNewCryptoWriter(t *testing.T) {
	key := "123456"
	memfile := new(bytes.Buffer)//创建一个内存文件
	w := NewCryptoWriter(memfile,key)
	w.Write([]byte("hello"))

	r := NewCryptoReader(memfile, key)
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)

	if string(buf[:n]) != "hello"{
		t.Error("not equal: %s, %s", buf[:n], "hello")
	}else{
		fmt.Println("单元测试通过")
	}
}

/*
基准测试 Benchmark开头 测试性能
 */
 func BenchmarkNewCryptoWriter(b *testing.B) {
	 buf := []byte(strings.Repeat("a", 1024))
	 w := NewCryptoWriter(ioutil.Discard, "123456")//Discard 把加密后的数据扔掉
	 for i := 0; i < b.N; i ++{
	 	n, _ := w.Write(buf)
	 	b.SetBytes(int64(n))//记录吞吐量
	 }
 }
 // 吞吐量跟时延是对立的指标 吞吐量大 时延就会上升
 /*
 goos: darwin
goarch: amd64
pkg: demo11/mycrypto
500000	      2007 ns/op	 510.16 MB/s
PASS
  */
// 2007 ns/op 单次循环时间
// 循环次数一定是 b.N
// buf 越小越快  但是吞吐量也会下降