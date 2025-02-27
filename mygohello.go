package main

import (
	"fmt"
	"net/http"
)

func additionHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	aStr := values.Get("name")
	fmt.Printf("My Hello, %s!\n", aStr)
	hello(aStr)
}

func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	http.HandleFunc("/hello", additionHandler)

	fmt.Println("Server listening on port 8080....")
	http.ListenAndServe(":8080", nil)
}