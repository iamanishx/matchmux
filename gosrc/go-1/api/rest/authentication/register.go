package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ipc/go-1/core/services"
	"ipc/go-1/core/models"
)



func Register(w http.ResponseWriter, r *http.Request) {

	var creds models.Credentials
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid JSON format: %v", err)
		return
	}

	if creds.Email == "" || creds.Phone == "" || creds.Name == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing required fields")
		return
	}
     
	_ , err := services.CreateUser(&creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating user: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")

}
