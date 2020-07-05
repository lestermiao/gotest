package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/lestermiao/gotest/chat"
)

func main(){
	
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	log.Printf("---2--")
	if err!= nil {
		log.Fatalf("cannot connect %v", err)
	}
		c := chat.NewChatServiceClient(conn)

		message := chat.Message{
			Body: "ha ha from  cleiint",
		}

		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("error calling hello %v", err)
		}

		log.Printf("Response : %v", response.Body)
	
}