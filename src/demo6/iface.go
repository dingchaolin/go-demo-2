package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

func handle( w http.ResponseWriter, r *http.Request ){
	fmt.Fprintf( w, "hello\n")
}
func main(){
	f, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Fprintf( f, "hello\n")
}