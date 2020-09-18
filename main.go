package main

import (
	"log"
	"net"
	"time"

	pb "github.com/salaleser/scraper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
	pb.ScraperServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)
	pb.RegisterScraperServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
