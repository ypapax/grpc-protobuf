package main

import (
	"context"
	"github.com/ypapax/grpc-protobuf/chat"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Hello from the client!",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when calling sayHello: %+v", err)
	}

	log.Printf("Response from Server: %+v", response.Body)
}
