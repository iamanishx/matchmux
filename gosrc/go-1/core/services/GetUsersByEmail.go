package services

import (
	"context"
	"ipc/db"
	"ipc/ent"
	"ipc/ent/users"
)

func GetUserByEmail(email string) (*ent.Users, error) {
	client := db.EntClient()
	defer client.Close()
	ctx := context.Background()
	user, err := client.Users.Query().
		Where(
			users.EmailEQ(email),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return user , nil

}
