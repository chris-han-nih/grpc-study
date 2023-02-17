package main

import (
	"context"
	"github.com/chris-han-nih/model/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var conn, err = grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	c := account.NewAccountServiceClient(conn)
	getAccount, err := c.GetAccount(context.Background(), &account.GetAccountRequest{Id: 1})
	if err != nil {
		return
	}
	log.Printf("GetAccount: %v", getAccount)
}
