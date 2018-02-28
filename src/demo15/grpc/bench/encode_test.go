package main

import (
	"encoding/json"
	"testing"

	"demo15/grpc/myproto"
	"github.com/golang/protobuf/proto"
)

func BenchmarkProto(b *testing.B) {
	var p myproto.Person
	p.Id = 1
	p.Name = "dcl"
	p.Email = "dcl@xx.com"
	p.Phones = []*myproto.PhoneNumber{
		{Number: "123456", Type: myproto.PhoneType_MOBILE},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&p)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	var p myproto.Person
	p.Id = 1
	p.Name = "dcl"
	p.Email = "dcl@xx.com"
	p.Phones = []*myproto.PhoneNumber{
		{Number: "123456", Type: myproto.PhoneType_MOBILE},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&p)
		if err != nil {
			panic(err)
		}
	}
}
