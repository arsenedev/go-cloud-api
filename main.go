package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go Cloud API!")
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bonjour from Go Cloud API!")

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
	println("Server is running on http://localhost:8080")
}
