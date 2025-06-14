package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Credentials struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
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
	
	fmt.Println("Received credentials:", creds)

}
