package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Stu1 struct{
	Id int
	Name string
}

func main(){
	str := `{"Id":2, "Name":"dcl"}`
	var s Stu1
	err := json.Unmarshal([]byte(str), &s)

	if err != nil {
		log.Fatal("unmaschal error: %s", err )
	}

	fmt.Println( s )
}
