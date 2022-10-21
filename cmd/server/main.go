package main

import (
	"context"
	"log"
	"net"

	"github.com/dannielss/go-grpc/pb"
	"google.golang.org/grpc"
)

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

func (s *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := ":5000"

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	grpcError := grpcServer.Serve(listener)

	if grpcError != nil {
		log.Fatal(grpcError)
	}
}
