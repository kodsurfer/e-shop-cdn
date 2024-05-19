package auth

import (
	"context"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/auth"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/auth/pb"
	"slices"
)

var _ auth.IClient = (*Client)(nil)

// UserTokenMock represent simple mock for testing
type UserTokenMock struct {
	Token string
	User  *pb.UserData
}

var mockedUser = &UserTokenMock{
	Token: "1",
	User: &pb.UserData{
		Id:        "664467965d460726d18e9103",
		Phone:     "770812312332",
		Email:     "test@mail.ru",
		FirstName: "Test",
		LastName:  "Test",
		IsActive:  true,
	},
}

// Client represent gRPC e-shop-auth client mock
type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Connect() error {
	return nil
}

func (c Client) Validate(ctx context.Context, token string) (*pb.UserData, error) {
	if token == mockedUser.Token {
		return mockedUser.User, nil
	}

	return nil, auth.ErrUserNotFound
}

func (c Client) FindUsersByIds(ctx context.Context, ids []string) ([]*pb.UserData, error) {
	if slices.Contains(ids, mockedUser.User.Id) {
		return []*pb.UserData{mockedUser.User}, nil
	}

	return []*pb.UserData{}, nil
}
