package account

import "context"

type Server struct {
	UnimplementedAccountServiceServer
}

func (s *Server) GetAccount(ctx context.Context, in *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, nil
}

func (s *Server) CreateAccount(ctx context.Context, in *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, nil
}
