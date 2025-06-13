package server

import (
	"fmt"
	"net/http"
)

func Server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the Go server!")
	})

	address := ":5000"
	fmt.Printf("Starting server on http://localhost%s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic("Error starting the server :" + err.Error())
	}
}
