package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Simple Web Framework")

	r := &router{make(map[string]map[string]http.HandlerFunc)}

	r.HandleFunc("GET", "/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "This is an index page.")
	})

	r.HandleFunc("GET", "/about", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "This is an about page.")
		fmt.Println("/about call")
	})

	r.HandleFunc("GET", "/users/:id", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Retrieve user")
		fmt.Println("Retrieve user")
	})

	r.HandleFunc("POST", "/users", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Create user")
		fmt.Println("Create user")
	})

	http.ListenAndServe(":8080", r)
}
