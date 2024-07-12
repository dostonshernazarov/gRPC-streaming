// server.go
package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "lessons/rpc-streaming/server-streaming/server/genproto/numbers"
)

type server struct {
	pb.UnimplementedNumberServiceServer
}

func (s *server) GetNumbers(req *pb.NumberRequest, stream pb.NumberService_GetNumbersServer) error {
	max := req.Max
	for i := int32(1); i <= max; i++ {
		if err := stream.Send(&pb.NumberResponse{Number: i}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNumberServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
