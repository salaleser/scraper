package main

import (
	"context"
	"log"

	pb "github.com/salaleser/scraper/scraper"
)

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Printf("Received: %v\n", in.Ok)

	return &pb.PingReply{
		Ok: "ok",
	}, nil
}
