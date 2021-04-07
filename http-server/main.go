package main

import (
	"fmt"
	"net/http"
)

func main() {
	const port = "3000"
	fmt.Println("Listening on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the server!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":"+port, nil)
}
