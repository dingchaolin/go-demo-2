package main

import (
	"time"
	"log"
)

func indexName()string{
	date := time.Now().Format("20060102")
	return "topic-" + date
}

func main(){
	log.Print( indexName())
}