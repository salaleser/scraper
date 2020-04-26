package scraper

import (
	"encoding/json"
	"log"
	"strings"
)

// Metadata contains application's metadata
type Metadata struct { // TODO add more fields
	Title       string
	AppID       string
	ArtistName  string
	Rating      string
	ReleaseDate string
	Subtitle    string
	Description string
	Screenshot1 string // TODO add array
	Logo        string
}

type StoryResponse struct {
	Canvas           Canvas
	Label            string
	ID               string
	CardIds          []string
	RelatedContent   map[string]Result
	EditorialArtwork EditorialArtwork
	Kind             string
	Link             Link
	DisplayStyle     string
	EditorialNotes   EditorialNotes
	CardDisplayStyle string
	DisplaySubStyle  string
}

func parseAsIDsBody(body []byte) []Metadata {
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as IDs (234): %q", err.Error())
		return []Metadata{} // TODO handle error
	}

	metadatas := make([]Metadata, 0)
	for _, result := range data.StorePlatformData.NativeSearchLockup.Results {
		if result.Kind != "iosSoftware" {
			continue
		}

		metadata := Metadata{
			Title: result.Name,
			AppID: result.ID,
		}

		metadatas = append(metadatas, metadata)
	}

	return metadatas
}

func parseAsMetadataBody(body []byte) Metadata {
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as metadata (1): %q", err.Error())
		return Metadata{} // TODO handle error
	}

	var metadata Metadata
	for _, result := range data.StorePlatformData.ProductDv.Results {
		var screenshot1 string
		for _, screenshots := range result.ScreenshotsByType {
			if len(screenshots) == 0 {
				continue
			}

			screenshot1 = strings.Replace(screenshots[0].URL, "{w}x{h}bb.{f}", "512x512bb.png", -1)
		}

		metadata = Metadata{
			AppID:       result.ID,
			ArtistName:  result.ArtistName,
			Rating:      result.UserRating.AriaLabelForRatings,
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

func parseGpIDsBody(body []byte) []Metadata {
	var data1 [][]interface{}
	if err := json.Unmarshal(body, &data1); err != nil {
		log.Printf("Error while trying to unmarshal gp IDs (1): %q", err.Error())
		return []Metadata{} // TODO handle error
	}

	d := data1[0]

	if d[0] != "wrb.fr" {
		log.Printf("The first metadata section element isn't \"wrb.fr\" (%q).", d[0])
		return []Metadata{} // TODO handle error
	}

	if d[1] != "lGYRle" {
		log.Printf("The second metadata section element isn't \"lGYRle\" (%q).", d[0])
		return []Metadata{} // TODO handle error
	}

	if d[2] == nil {
		log.Printf("Error while parsing (386).")
		return []Metadata{} // TODO handle error
	}

	var data2 []interface{}
	if err := json.Unmarshal([]byte(d[2].(string)), &data2); err != nil {
		log.Printf("Error while trying to unmarshal gp IDs (2): %q", err.Error())
		return []Metadata{} // TODO handle error
	}

	// FIXME
	if len(data2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[0].([]interface{})) < 2 {
		return []Metadata{}
	}

	// FIXME может выйти за пределы массива
	metadatas := make([]Metadata, 5)
	for i := 0; i < 5; i++ {
		metadata := Metadata{
			Title: data2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[0].([]interface{})[i].([]interface{})[2].(string),
			AppID: data2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[0].([]interface{})[i].([]interface{})[12].([]interface{})[0].(string),
		}

		metadatas = append(metadatas, metadata)
	}

	return metadatas
}

func parseGpMetadataBody(body []byte) Metadata {
	var data1 [][]interface{}
	if err := json.Unmarshal(body, &data1); err != nil {
		log.Printf("Error while trying to unmarshal gp metadata (1): %q", err.Error())
		return Metadata{} // TODO handle error
	}

	d := data1[0]

	if d[0] != "wrb.fr" {
		log.Printf("The first metadata section element isn't \"wrb.fr\" (%q).", d[0])
		return Metadata{} // TODO handle error
	}

	if d[1] != "jLZZ2e" {
		log.Printf("The second metadata section element isn't \"jLZZ2e\" (%q).", d[0])
		return Metadata{} // TODO handle error
	}

	if d[2] == nil {
		log.Printf("Error while parsing (567).")
		return Metadata{} // TODO handle error
	}

	var data2 [][][]interface{}
	if err := json.Unmarshal([]byte(d[2].(string)), &data2); err != nil {
		log.Printf("Error while trying to unmarshal gp metadata (2): %q", err.Error())
		return Metadata{} // TODO handle error
	}

	return Metadata{
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
	for _, v := range data.StorePlatformData.EditorialItemProduct.Results {
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
