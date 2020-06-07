package scraper

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	as "github.com/salaleser/appstoreapi"
	pb "github.com/salaleser/scraper/scraper"
	"google.golang.org/grpc"
)

type server struct {
	pb.ScraperServer
}

func (s *server) Room(ctx context.Context, in *pb.RoomRequest) (*pb.RoomReply, error) {
	log.Printf("Received: %v", in.GetId())
	data, err := as.Room(in.GetId(), in.Country, in.Language)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}

	return &pb.RoomReply{
		Data: &pb.Data{
			PageData: &pb.PageData{
				AdamId: uint32(data.PageData.AdamID),
			},
		},
	}, nil
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
