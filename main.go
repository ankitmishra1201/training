package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", HelloWrold)
	http.ListenAndServe(":50060", nil)
}

func HelloWrold(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
