package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "server/pb"

	"google.golang.org/grpc"
)

const (
	port = ":8181"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, &server{})

	fmt.Printf("server started...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
