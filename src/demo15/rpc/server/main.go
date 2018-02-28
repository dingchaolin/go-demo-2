package main

import (
	"net/rpc"
	"net"
	"log"
	"demo15/rpc/common"
)

type MathService struct{

}


func (m *MathService) Add(request *common.AddRequest, reply *common.AddResponse) error{
	reply.Result = request.M + request.N
	return nil
}

func (m *MathService) Mul(request *common.AddRequest, reply *common.AddResponse)error{
	reply.Result = request.M * request.N
	return nil
}

func main(){
	mathService := new(MathService)
	rpc.Register(mathService)

	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}

	rpc.Accept(l)

}