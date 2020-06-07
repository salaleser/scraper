package scraper

import (
	"context"
	"log"
	"net"

	pb "github.com/salaleser/scraper/scraper"
	"google.golang.org/grpc"
)

type server struct {
	pb.ScraperServer
}

func (s *server) Room(ctx context.Context, in *pb.RoomRequest) (*pb.RoomReply, error) {
	log.Printf("Received: %v", in.GetId())
	// _, _ = as.Room(in.GetId(), in.Country, in.Language)

	return &pb.RoomReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScraperServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
