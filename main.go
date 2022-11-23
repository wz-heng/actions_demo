package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	log.Print("hello")
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}
