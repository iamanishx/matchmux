package services

import (
	"context"
	"ipc/db"
	"ipc/go-1/core/models"
)

func CreateUser(c *models.Credentials) (string, error) {
	ctx := context.Background()
	client := db.EntClient()
	defer client.Close()
	_, err := client.Users.Create().
		SetEmail(c.Email).
		SetPhone(c.Phone).
		SetName(c.Name).
		SetPassword(c.Password).
		Save(ctx)
	if err != nil {
		return "", err
	}
	return "User created successfully", nil
}
