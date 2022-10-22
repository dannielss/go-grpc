package main

import (
	"context"
	"log"
	"net"

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
