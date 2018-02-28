package main

import (
	"demo15/grpc/myproto"
	"io/ioutil"
	"os"
	"log"
	"fmt"
	"github.com/golang/protobuf/proto"
)
// protobuf 数据比较紧凑
func main(){
	var p myproto.Person
	buf, _  := ioutil.ReadAll(os.Stdin)
	err := proto.Unmarshal(buf, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( p.String())

}