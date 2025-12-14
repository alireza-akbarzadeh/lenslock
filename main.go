package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello, World!</h1>")
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}
