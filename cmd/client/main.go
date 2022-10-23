package main

import (
	"context"
	"log"

	"github.com/dannielss/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello GRPC",
	}

	res, err := client.RequestMessage(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res.GetStatus())
}
