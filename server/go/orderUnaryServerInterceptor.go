package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("============= [Server Interceptor] ", info.FullMethod)

	m, err := handler(ctx, req)

	resp, ok := m.(string)
	if ok {
		log.Printf(" Post Proc Message: %s", resp)
	}

	return m, err
}
