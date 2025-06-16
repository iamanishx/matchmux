package authentication

import (
	"encoding/json"
	"ipc/go-1/core/models"
	"ipc/go-1/core/services"
	"net/http"
)

func Verify(w http.ResponseWriter, r *http.Request) {

	var otp models.Otpverification
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&otp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if otp.Code == "" || otp.UserID.String() == "00000000-0000-0000-0000-000000000000" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	response, err := services.VerifyOtp(&otp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Error verifying OTP: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
