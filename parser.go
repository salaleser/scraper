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

type FeaturedStory struct {
	Canvas           Canvas           `json:"canvas"`
	Label            string           `json:"label"`
	ID               string           `json:"id"`
	CardIds          []string         `json:"cardIds"`
	RelatedContent   RelatedContent   `json:"relatedContent"`
	EditorialArtwork EditorialArtwork `json:"editorialArtwork"`
	Kind             string           `json:"kind"`
	Link             Link             `json:"link"`
	DisplayStyle     string           `json:"displayStyle"`
	EditorialNotes   EditorialNotes   `json:"editorialNotes"`
	CardDisplayStyle string           `json:"cardDisplayStyle"`
	DisplaySubStyle  string           `json:"displaySubStyle"`
}

func parseAsIDsBody(body []byte) []Metadata {
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as IDs (234): %q", err.Error())
		return []Metadata{} // TODO handle error
	}

	results := data.StorePlatformData.NativeSearchLockup.Results.(map[string]interface{})

	metadatas := make([]Metadata, 0)
	for _, v := range results {
		data, err := json.Marshal(v)
		if err != nil {
			log.Printf("Error while trying to marshal as IDs (235): %q", err.Error())
			return []Metadata{} // TODO handle error
		}

		var result Result
		if err := json.Unmarshal(data, &result); err != nil {
			log.Printf("Error while trying to unmarshal as IDs (234): %q", err.Error())
			return []Metadata{} // TODO handle error
		}

		if result.Kind == "iosSoftware" {
			metadatas := make([]Metadata, 5)
			for i := 0; i < 5; i++ {
				metadatas[i] = Metadata{
					Title: result.Name,
					AppID: result.ID,
				}
			}
		}
	}

	return metadatas
}

func parseAsMetadataBody(body []byte) Metadata {
	var data Page
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as metadata (1): %q", err.Error())
		return Metadata{} // TODO handle error
	}

	results := data.StorePlatformData.ProductDv.Results.(map[string]interface{})

	var metadata Metadata
	for _, v := range results {
		data, err := json.Marshal(v)
		if err != nil {
			log.Printf("Error while trying to marshal as IDs (235): %q", err.Error())
			return Metadata{} // TODO handle error
		}

		var result Result
		if err := json.Unmarshal(data, &result); err != nil {
			log.Printf("Error while trying to unmarshal as IDs (234): %q", err.Error())
			return Metadata{} // TODO handle error
		}

		screenshotsData, err := json.Marshal(result.ScreenshotsByType)
		if err != nil {
			log.Printf("Error while trying to marshal as IDs (235): %q", err.Error())
			return Metadata{} // TODO handle error
		}

		var screenshotsByTypeElement ScreenshotsByTypeElement
		if err := json.Unmarshal(screenshotsData, &screenshotsByTypeElement); err != nil {
			log.Printf("Error while trying to unmarshal as IDs (234): %q", err.Error())
			return Metadata{} // TODO handle error
		}

		var screenshot1 string
		for _, screenshots := range screenshotsByTypeElement.Screenshots {
			screenshot1 = strings.Replace(screenshots.URL, "{w}x{h}bb.{f}", "512x512bb.png", -1)
			break
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

	metadatas := make([]Metadata, 5)
	for i := 0; i < 5; i++ {
		metadatas[i] = Metadata{
			Title: data2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[0].([]interface{})[i].([]interface{})[2].(string),
			AppID: data2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[0].([]interface{})[i].([]interface{})[12].([]interface{})[0].(string),
		}
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
		// AppID:       result.ID,
		// ArtistName:  result.ArtistName,
		// Rating:      result.UserRating.AriaLabelForRatings,
		// ReleaseDate: result.ReleaseDate,
		// Title:       result.Name,
		// Subtitle:    result.Subtitle,
		// Description: result.Description.Standard,
		// Screenshot1: screenshot1,
		// Logo:        strings.Replace(result.Artwork.URL, "{w}x{h}bb.{f}", "128x128bb.png", -1),
		Title:       data2[0][0][0].(string),
		Subtitle:    data2[0][10][1].([]interface{})[1].(string),
		Description: data2[0][10][0].([]interface{})[1].(string),
		Screenshot1: data2[0][12][0].([]interface{})[0].([]interface{})[3].([]interface{})[2].(string),
		Logo:        data2[0][12][1].([]interface{})[3].([]interface{})[2].(string),
	}
}

func parseAsStory(body []byte) FeaturedStory {
	var data Story
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal as story (1): %q", err.Error())
		return FeaturedStory{} // TODO handle error
	}

	story := FeaturedStory{}

	return story
}
