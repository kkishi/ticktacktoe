package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := tpb.NewTickTackToeClient(conn)

	stream, err := client.Game(context.Background())
	if err != nil {
		log.Fatalf("client.Game failed: %v", err)
	}
	for {
		req := &tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: "Tick",
				},
			},
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("stream.Send failed: %v\n", err)
		}
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("stream.Recv failed: %v\n", err)
		}
		log.Printf("res: %v\n", res)
	}
}
