package graph

import (
	"context"
	"test-dikurium/graph/model"
)

func GetUserFromCtx(ctx context.Context, email string) (*model.User, error) {
	load, err := GetUserLoader(ctx).Load(email)
	if err != nil {
		return nil, err
	}
	var user model.User
	user.Username = load.Username
	user.Email = load.Email
	user.Userid = load.Userid
	return &user, nil
}
