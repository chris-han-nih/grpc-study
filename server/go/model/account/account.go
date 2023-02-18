package account

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	UnimplementedAccountServiceServer
}

func (s *Server) GetAccount(ctx context.Context, in *GetAccountRequest) (*GetAccountResponse, error) {
	account := &Account{
		Id:        in.Id,
		Name:      "Chris",
		Email:     "namilhan@gmail.com",
		Phone:     "010-1234-5678",
		Avatar:    "https://avatars0.githubusercontent.com/u/16812446?s=460&v=4",
		CreatedAt: timestamppb.New(time.Now()),
		Address: &Address{
			Street: "gangnam",
			City:   "seoul",
			State:  "korea",
			Zip:    "12345",
		},
	}

	return &GetAccountResponse{Account: account}, nil
}

func (s *Server) CreateAccount(ctx context.Context, in *CreateAccountRequest) (*CreateAccountResponse, error) {
	// TODO: Implement the account creation logic
	return nil, nil
}
