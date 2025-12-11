package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HELLO WORD!")
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Sadia", "I am a student", "I am Learning...")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", abouthandler)

	fmt.Println("Server running on: 3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error staring the server", err)
	}
}
