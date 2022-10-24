package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/dannielss/go-grpc/pb"
	"google.golang.org/grpc"
)

const PORT = ":5000"

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (s *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("message received: ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}

func (s *Server) GetUsers(req *pb.Request, u pb.SendMessage_GetUsersServer) error {
	log.Print("message received: ", req.GetMessage())

	users := []*pb.User{
		{Name: "Daniel", Age: 21},
		{Name: "Daniela", Age: 22},
	}

	// sequentially

	// for _, user := range users {
	// 	resp := pb.User{Name: user.Name, Age: user.Age}

	// 	if err := u.Send(&resp); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// concurrently

	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go func(user *pb.User) {
			defer wg.Done()
			resp := pb.User{Name: user.GetName(), Age: user.GetAge()}
			if err := u.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}(user)
	}

	wg.Wait()

	return nil
}

func (s *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatal(err)
	}

	grpcError := grpcServer.Serve(listener)

	if grpcError != nil {
		log.Fatal(grpcError)
	}
}
