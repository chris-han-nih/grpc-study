package main

import (
	"github.com/chris-han-nih/model/account"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50052"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	account.RegisterAccountServiceServer(s, &account.Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
