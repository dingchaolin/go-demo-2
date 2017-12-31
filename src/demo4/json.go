package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Stu struct{
	Id int
	Name string
}

func main(){
	s := Stu{
		Id:2,
		Name:"dcl",
	}

	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal("marshal err: %s", err )
	}
	fmt.Println( string(buf) )
}
