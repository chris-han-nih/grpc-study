package main

import (
	"fmt"
	"github.com/chris-han-nih/model/account"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50052"
)

func main() {
	if err := runServer(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func runServer() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	s := grpc.NewServer()
	account.RegisterAccountServiceServer(s, &account.Server{})
	if err = s.Serve(lis); err != nil {
		return fmt.Errorf("failed to server: %v", err)
	}
	return nil
}
