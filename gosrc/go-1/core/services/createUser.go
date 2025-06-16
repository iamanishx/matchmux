package services

import (
	"context"
	"ipc/db"
	"ipc/go-1/core/helper"
	"ipc/go-1/core/models"
	"ipc/go-1/mail"
	"time"
)

func CreateUser(c *models.Credentials) (map[string]interface{}, error) {
	ctx := context.Background()
	client := db.EntClient()
	defer client.Close()
	user, err := client.Users.Create().
		SetEmail(c.Email).
		SetPhone(c.Phone).
		SetName(c.Name).
		SetPassword(c.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	otp := helper.GenerateOTP()

	_, err = client.Otp.Create().
		SetCode(otp).
		SetUserID(user.ID).
		SetExpiresAt(time.Now().Add(10 * time.Minute)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	mail.SendEmail(c.Email, otp)

response := map[string]interface{}{
    "status": "pending_verification",
    "message": "Please check your email for verification code",
    "user_id": user.ID, 
}
	return response, nil
}
