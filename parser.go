package scraper

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

// MetadataResponse is a vitalina's Application's metadata structure.
type MetadataResponse struct { // TODO add more fields
	Title       string
	Link        string
	AppID       string
	ArtistName  string
	Rating      float32
	ReleaseDate string
	Subtitle    string
	Description string
	Screenshot1 string // TODO add array
	Logo        string
}

// StoryResponse is a vitalina's Story structure.
type StoryResponse struct {
	Canvas           Canvas
	Label            string
	ID               string
	CardIds          []string
	RelatedContent   map[string]Result
	EditorialArtwork map[string]Artwork
	Kind             string
	Link             Link
	DisplayStyle     string
	EditorialNotes   EditorialNotes
	CardDisplayStyle string
	DisplaySubStyle  string
}

// RoomResponse is a vitalina's Room structure.
type RoomResponse struct {
	Title       string
	Link        string
	AppID       string
	ArtistName  string
	Rating      float32
	ReleaseDate string
	Subtitle    string
	Description string
	Screenshot1 string // TODO add array
	Logo        string
}

func parseAsIDs(body []byte) []MetadataResponse {
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as IDs (234): %q", err.Error())
		return []MetadataResponse{} // TODO handle error
	}

	metadatas := make([]MetadataResponse, 0)
	for _, result := range data.StorePlatformData["native-search-lockup"].Results {
		if result.Kind != "iosSoftware" {
			continue
		}

		metadata := MetadataResponse{
			Title:  result.Name,
			AppID:  result.ID,
			Rating: result.UserRating.Value,
		}

		metadatas = append(metadatas, metadata)
	}

	return metadatas
}

func parseAsMetadata(body []byte) MetadataResponse {
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as metadata (1): %q", err.Error())
		return MetadataResponse{} // TODO handle error
	}

	var metadata MetadataResponse
	for _, result := range data.StorePlatformData["product-dv"].Results {
		var screenshot1 string
		for _, screenshots := range result.ScreenshotsByType {
			if len(screenshots) == 0 {
				continue
			}

			screenshot1 = strings.Replace(screenshots[0].URL, "{w}x{h}bb.{f}", "512x512bb.png", -1)
		}

		metadata = MetadataResponse{
			AppID:       result.ID,
			Link:        result.Link.URL,
			ArtistName:  result.ArtistName,
			Rating:      result.UserRating.Value,
			ReleaseDate: result.ReleaseDate,
			Title:       result.Name,
			Subtitle:    result.Subtitle,
			Description: result.Description.Standard,
			Screenshot1: screenshot1,
			Logo:        strings.Replace(result.Artwork.URL, "{w}x{h}bb.{f}", "128x128bb.png", -1),
		}
	}

	return metadata
}

// ParsePage parses App Store root page and returns structure
func ParsePage(body []byte) (Page, error) {
	var page Page
	err := json.Unmarshal(body, &page)
	if err != nil {
		return Page{}, err
	}

	return page, nil
}

func parseAsRoom(body []byte) RoomResponse {
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as room (1): %q", err.Error())
		return RoomResponse{} // TODO handle error
	}

	var room RoomResponse
	for _, result := range data.StorePlatformData["lockup"].Results {
		var screenshot1 string
		for _, screenshots := range result.ScreenshotsByType {
			if len(screenshots) == 0 {
				continue
			}

			screenshot1 = strings.Replace(screenshots[0].URL, "{w}x{h}bb.{f}", "512x512bb.png", -1)
		}

		room = RoomResponse{
			AppID:       result.ID,
			Link:        result.Link.URL,
			ArtistName:  result.ArtistName,
			Rating:      result.UserRating.Value,
			ReleaseDate: result.ReleaseDate,
			Title:       result.Name,
			Subtitle:    result.Subtitle,
			Description: result.Description.Standard,
			Screenshot1: screenshot1,
			Logo:        strings.Replace(result.Artwork.URL, "{w}x{h}bb.{f}", "128x128bb.png", -1),
		}
	}

	return room
}

func parseGpIDs(body []byte) []MetadataResponse {
	var data1 [][]interface{}
	if err := json.Unmarshal(body, &data1); err != nil {
		log.Printf("Error while trying to unmarshal gp IDs (1): %q", err.Error())
		return []MetadataResponse{} // TODO handle error
	}

	d := data1[0]

	if d[0] != "wrb.fr" {
		log.Printf("The first metadata section element isn't \"wrb.fr\" (%q).", d[0])
		return []MetadataResponse{} // TODO handle error
	}

	if d[1] != "lGYRle" {
		log.Printf("The second metadata section element isn't \"lGYRle\" (%q).", d[0])
		return []MetadataResponse{} // TODO handle error
	}

	if d[2] == nil {
		log.Printf("Error while parsing (386).")
		return []MetadataResponse{} // TODO handle error
	}

	var data2 []interface{}
	if err := json.Unmarshal([]byte(d[2].(string)), &data2); err != nil {
		log.Printf("Error while trying to unmarshal gp IDs (2): %q", err.Error())
		return []MetadataResponse{} // TODO handle error
	}

	i1 := data2[0].([]interface{})
	if i1 == nil {
		i1JSON, _ := json.Marshal(data2[0])
		log.Printf("cast interface 1: %q", errors.New(string(i1JSON)))
		return []MetadataResponse{} // TODO handle error
	}

	i1_1 := i1[1]
	if i1_1 == nil {
		i1_1JSON, _ := json.Marshal(i1)
		log.Printf("cast interface 1.1: %q", errors.New(string(i1_1JSON)))
		return []MetadataResponse{} // TODO handle error
	}

	i2 := i1_1.([]interface{})
	if i2 == nil {
		i2JSON, _ := json.Marshal(i1_1)
		log.Printf("cast interface 2: %q", errors.New(string(i2JSON)))
		return []MetadataResponse{} // TODO handle error
	}

	i3 := i2[0].([]interface{})
	if i3 == nil {
		i3JSON, _ := json.Marshal(i2)
		log.Printf("cast interface 3: %q", errors.New(string(i3JSON)))
		return []MetadataResponse{} // TODO handle error
	}

	i4 := i3[0].([]interface{})
	if i4 == nil {
		i4JSON, _ := json.Marshal(i3)
		log.Printf("cast interface 4: %q", errors.New(string(i4JSON)))
		return []MetadataResponse{} // TODO handle error
	}

	i5 := i4[0].([]interface{})
	if i5 == nil {
		i5JSON, _ := json.Marshal(i4)
		log.Printf("cast interface 5: %q", errors.New(string(i5JSON)))
		return []MetadataResponse{} // TODO handle error
	}

	// FIXME
	if len(i5) < 2 {
		log.Printf("len gp json array check: %q", errors.New("len < 2"))
		return []MetadataResponse{}
	}

	// FIXME interfaces
	metadatas := make([]MetadataResponse, 0)
	for _, d := range i5 {
		metadata := MetadataResponse{
			Title:  d.([]interface{})[2].(string),
			AppID:  d.([]interface{})[12].([]interface{})[0].(string),
			Rating: -1,
		}

		metadatas = append(metadatas, metadata)
	}

	return metadatas
}

func parseGpMetadata(body []byte) MetadataResponse {
	var data1 [][]interface{}
	if err := json.Unmarshal(body, &data1); err != nil {
		log.Printf("Error while trying to unmarshal gp metadata (1): %q", err.Error())
		return MetadataResponse{} // TODO handle error
	}

	d := data1[0]

	if d[0] != "wrb.fr" {
		log.Printf("The first metadata section element isn't \"wrb.fr\" (%q).", d[0])
		return MetadataResponse{} // TODO handle error
	}

	if d[1] != "jLZZ2e" {
		log.Printf("The second metadata section element isn't \"jLZZ2e\" (%q).", d[0])
		return MetadataResponse{} // TODO handle error
	}

	if d[2] == nil {
		log.Printf("Error while parsing (567).")
		return MetadataResponse{} // TODO handle error
	}

	var data2 [][][]interface{}
	if err := json.Unmarshal([]byte(d[2].(string)), &data2); err != nil {
		log.Printf("Error while trying to unmarshal gp metadata (2): %q", err.Error())
		return MetadataResponse{} // TODO handle error
	}

	return MetadataResponse{
		// AppID: appID,
		ArtistName: data2[0][12][5].([]interface{})[1].(string),
		// ReleaseDate: data2[0][6][0][1].(float32),
		// Rating:      data2[0][0][0].(string),
		Title:       data2[0][0][0].(string),
		Subtitle:    data2[0][10][1].([]interface{})[1].(string),
		Description: data2[0][10][0].([]interface{})[1].(string),
		Screenshot1: data2[0][12][0].([]interface{})[0].([]interface{})[3].([]interface{})[2].(string),
		Logo:        data2[0][12][1].([]interface{})[3].([]interface{})[2].(string),
	}
}

func parseAsStory(body []byte) StoryResponse {
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as story (1): %q", err.Error())
		return StoryResponse{} // TODO handle error
	}

	var result Result
	for _, v := range data.StorePlatformData["editorial-item-product"].Results {
		result = v
		break
	}

	return StoryResponse{
		Canvas:           result.Canvas,
		Label:            result.Label,
		ID:               result.ID,
		CardIds:          result.CardIds,
		RelatedContent:   result.RelatedContent,
		EditorialArtwork: result.EditorialArtwork,
		Kind:             result.Kind,
		Link:             result.Link,
		DisplayStyle:     result.DisplayStyle,
		EditorialNotes:   result.EditorialNotes,
		CardDisplayStyle: result.CardDisplayStyle,
		DisplaySubStyle:  result.DisplaySubStyle,
	}
}
