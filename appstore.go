package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	api "github.com/salaleser/appstoreapi"
	pb "github.com/salaleser/scraper/scraper"
)

func (s *server) AppStoreRoom(ctx context.Context, in *pb.AppStoreRoomRequest) (*pb.AppStoreRoomReply, error) {
	log.Printf("Received: %v\n", in.GetId())
	data, err := api.Room(in.GetId(), in.CountryCode, in.Language)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scraper: %v", err)
	}

	contentIDs := make([]uint32, 0)
	for _, id := range data.PageData.AdamIDs {
		contentIDs = append(contentIDs, uint32(id))
	}

	fcKind, err := strconv.Atoi(data.PageData.FcKind)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}

	storeFront, err := strconv.Atoi(data.PageData.MetricsBase.StoreFront)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}

	language, err := strconv.Atoi(data.PageData.MetricsBase.Language)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}

	return &pb.AppStoreRoomReply{
		Id:         uint32(data.PageData.AdamID),
		FcKind:     uint32(fcKind),
		StoreFront: uint32(storeFront),
		Language:   uint32(language),
		Title:      data.PageData.PageTitle,
		ContentIds: contentIDs,
	}, nil
}
