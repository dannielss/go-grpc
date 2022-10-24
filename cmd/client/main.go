package main

import (
	"context"
	"io"
	"log"
	"time"

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

	// Unary RPC
	// res, err := client.RequestMessage(context.Background(), req)

	// Server streaming RPC
	stream, err := client.GetUsers(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("Resp received: %+v", resp)
			time.Sleep(2 * time.Second)

		}
	}()

	<-done
	log.Printf("finished")
}
