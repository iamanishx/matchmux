package authentication

import (
	"context"
	"ipc/db"
)

func CreateUser(c *Credentials) (string, error) {
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
