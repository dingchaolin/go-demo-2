package main

import ("demo15/grpc/myproto"
"github.com/golang/protobuf/proto"
	"fmt"
)

func main(){
	var p myproto.Person
	p.Id = 1
	p.Name = "dcl"
	p.Email = "11234@qq.com"
	p.Phones = []*myproto.PhoneNumber{
		{Number:"123456", Type:myproto.PhoneType_MOBILE},
	}
	buf, err := proto.Marshal(&p)
	if err != nil {
		panic(err)
	}
	fmt.Println( string(buf))
}
