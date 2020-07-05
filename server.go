package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/lestermiao/gotest/chat"
)

func main(){
	lis, err := net.Listen("tcp", ":9000")
	if err != nil{
		log.Fatalf("Failed to port %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err:= grpcServer.Serve(lis); err!=nil{
		log.Fatalf("failed to start %v", err)
	}

}