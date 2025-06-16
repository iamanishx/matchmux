package restapi

import (
	"fmt"
	"ipc/go-1/api/rest/authentication"
	"net/http"
)

func StartRest() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the Go server!")
	})

	http.HandleFunc("/register", authentication.Register)
	http.HandleFunc("/verify", authentication.Verify)
	http.HandleFunc("/login", authentication.Login)

	address := ":5000"
	fmt.Printf("Starting server on http://localhost%s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic("Error starting the server :" + err.Error())
	}
}
