package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	fmt.Println("Server starting at :5000...")
	http.ListenAndServe(":5000", nil)
}
