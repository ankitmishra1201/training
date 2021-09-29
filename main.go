package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":50060", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
