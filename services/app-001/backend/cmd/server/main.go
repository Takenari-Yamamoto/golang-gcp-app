package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World Take!!!")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
