package services

import (
	"context"
	"ipc/db"
	"ipc/ent"
	"ipc/ent/otp"
	"ipc/ent/users"
	"ipc/go-1/core/models"
)

func VerifyOtp(c *models.Otpverification) (map[string]interface{}, error) {
	client := db.EntClient()
	defer client.Close()
	ctx := context.Background()
	otp, err := client.Otp.Query().
		Where(
			otp.CodeEQ(c.Code),
			otp.UserIDEQ(c.UserID),
		).
		Only(ctx)
	errorResponse := map[string]interface{}{
		"status":  "error",
		"message": "Invalid OTP or User ID",
	}
	if err != nil {
		if ent.IsNotFound(err) {
			return errorResponse, nil
		}
		return errorResponse, err
	}
	client.Users.Update().
		SetVerified(true).
		Where(
			users.IDEQ(c.UserID),
		).
		Exec(ctx)
	client.Otp.DeleteOne(otp).Exec(ctx)
	response := map[string]interface{}{
		"status":  "verified",
		"message": "OTP verified successfully please login",
	}

	return response, nil
}
