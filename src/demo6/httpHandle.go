package main

import (
	"net/http"
	"fmt"
	"time"
	"log"
)

type Monitor struct {
	counter int
}
/*
类似内部状态暴露
 */
func ( m *Monitor ) Run(){
	for{
		time.Sleep( time.Second )
		m.counter ++
	}
}

func ( m *Monitor ) ServeHTTP(w http.ResponseWriter,  r *http.Request ){
	fmt.Fprintf( w, "count:%d\n", m.counter )
}

func myHandle( w http.ResponseWriter, r *http.Request ){
	fmt.Fprintf(w, "hello world\n")
}

func main(){

	var m Monitor

	http.HandleFunc("/", myHandle)
	http.Handle("/monitor", &m)
	go m.Run()
	log.Fatal( http.ListenAndServe(":9090", nil ))

}

//curl http://localhost:9090/monitor