package main

import (
	"log"
	"net"

	"github.com/kkishi/ticktacktoe/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

const port = ":12345"

func main() {
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tpb.RegisterTickTackToeServer(s, &server.Impl{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
