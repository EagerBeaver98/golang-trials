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

	http.ListenAndServe(":"+port, nil)
}
