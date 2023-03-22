package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("============= [Server Interceptor] ", info.FullMethod)

	m, err := handler(ctx, req)

	log.Println("============= [Server Interceptor] ", m, err)
	if err == nil {
		log.Printf(" Post Proc Message: %s", m)
	}

	return m, err
}
