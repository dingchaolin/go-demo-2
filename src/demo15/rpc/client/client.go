package main

import (
	"net/rpc"
	"log"
	"demo15/rpc/common"
)
// google.golang.org/grpc
func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8021")
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	req := common.AddRequest{M: 20, N: 10}
	var reply common.AddResponse
	err = client.Call("MathService.Add", &req, &reply)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf( "result:%d", reply.Result)

}
