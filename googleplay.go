package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "github.com/salaleser/googleplayapi"
	pb "github.com/salaleser/scraper/proto"
)

func (s *server) GooglePlayApp(ctx context.Context, in *pb.GooglePlayAppRequest) (*pb.GooglePlayAppReply, error) {
	log.Printf("Received: %v\n", in.GetPackageName())
	data, err := api.App(in.GetPackageName(), in.GeoLocation, in.HumanLanguage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scraper gp app: %v", err)
		return &pb.GooglePlayAppReply{}, err
	}

	return &pb.GooglePlayAppReply{
		Rating:       data.Rating,
		StarsCount:   data.StarsCount,
		Stars_1Count: data.Stars1Count,
		Stars_2Count: data.Stars2Count,
		Stars_3Count: data.Stars3Count,
		Stars_4Count: data.Stars4Count,
		Stars_5Count: data.Stars5Count,
		PackageName:  data.AppID,
		Title:        data.Title,
	}, nil
}
