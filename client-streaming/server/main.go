package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "lessons/rpc-streaming/client-streaming/server/genproto/sum"
)

type server struct {
	pb.UnimplementedSumServiceServer
}

func (s *server) CalculateSum(stream pb.SumService_CalculateSumServer) error {
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.SumResponse{Sum: sum})
		}
		if err != nil {
			return err
		}
		sum += req.Number
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Server is running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
