package main

import (
	"net"
	"log"
	"bufio"
)

func handleConn( conn net.Conn ){
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')

	var content []byte
	conn.Write(content)

}

func main(){
	addr := ":8021"
	listenner, err := net.Listen("tcp", addr )
	if err != nil{
		log.Fatal(err)
	}

	defer listenner.Close()

	for{
		conn, err := listenner.Accept()
		if err != nil{
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}
