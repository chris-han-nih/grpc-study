package account

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	UnimplementedAccountServiceServer
}

func (s *Server) GetAccount(ctx context.Context, in *GetAccountRequest) (*GetAccountResponse, error) {
	return &GetAccountResponse{
		Account: &Account{
			Id:        1,
			Name:      "Chris",
			Email:     "namilhan@gmail.com",
			Phone:     "010-1234-5678",
			Avatar:    "https://avatars0.githubusercontent.com/u/16812446?s=460&v=4",
			CreatedAt: timestamppb.Now(),
			Address: &Address{
				Street: "gangnam",
				City:   "seoul",
				State:  "korea",
				Zip:    "12345",
			},
		},
	}, nil
}

func (s *Server) CreateAccount(ctx context.Context, in *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, nil
}
