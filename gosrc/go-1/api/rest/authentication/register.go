package authentication

import (
	"encoding/json"
	"fmt"
	"ipc/go-1/core/models"
	"ipc/go-1/core/services"
	"net/http"
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

	response, err := services.CreateUser(&creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating user: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	 if err := json.NewEncoder(w).Encode(response); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error encoding response")
        return
    }

}
