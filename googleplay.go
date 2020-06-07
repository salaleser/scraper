package main

import (
	"context"
	"log"

	api "github.com/salaleser/googleplayapi"
	pb "github.com/salaleser/scraper/scraper"
)

func (s *server) GooglePlayApp(ctx context.Context, in *pb.GooglePlayAppRequest) (*pb.GooglePlayAppReply, error) {
	log.Printf("Received: %v\n", in.GetPackageName())
	data := api.App(in.GetPackageName(), in.GeoLocation, in.HumanLanguage)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "scraper: %v", err)
	// }

	return &pb.GooglePlayAppReply{
		PackageName: data.AppID,
		Title:       data.Title,
	}, nil
}
