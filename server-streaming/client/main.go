// client.go
package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "lessons/rpc-streaming/server-streaming/client/genproto/numbers"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNumberServiceClient(conn)

	maxNum := int32(20)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(maxNum))
	defer cancel()

	stream, err := client.GetNumbers(ctx, &pb.NumberRequest{Max: maxNum})
	if err != nil {
		log.Fatalf("could not get numbers: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receiving: %v", err)
		}
		log.Printf("Received number: %d", resp.GetNumber())
	}
}
