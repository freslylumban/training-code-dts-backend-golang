package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("NAMA")
	port := os.Getenv("PORT")

	http.HandleFunc("/", hello)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	fmt.Println("Hello World, dari port", port, name)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
