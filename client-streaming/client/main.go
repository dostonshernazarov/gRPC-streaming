// client.go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "lessons/rpc-streaming/client-streaming/client/genproto/sum"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSumServiceClient(conn)

	stream, err := client.CalculateSum(context.Background())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}

	numbers := []int32{1, 2, 3, 4, 5}
	for _, number := range numbers {
		if err := stream.Send(&pb.SumRequest{Number: number}); err != nil {
			log.Fatalf("could not send number: %v", err)
		}
		println("Send number to server: ", number)
		time.Sleep(1 * time.Second)

	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}

	log.Printf("Sum: %d", resp.GetSum())
}
