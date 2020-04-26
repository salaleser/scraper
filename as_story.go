package scraper

type Story struct {
	Canvas           Canvas            `json:"canvas"`
	Label            string            `json:"label"`
	ID               string            `json:"id"`
	CardIds          []string          `json:"cardIds"`
	RelatedContent   map[string]Result `json:"relatedContent"`
	EditorialArtwork EditorialArtwork  `json:"editorialArtwork"`
	Kind             string            `json:"kind"`
	Link             Link              `json:"link"`
	DisplayStyle     string            `json:"displayStyle"`
	EditorialNotes   EditorialNotes    `json:"editorialNotes"`
	CardDisplayStyle string            `json:"cardDisplayStyle"`
	DisplaySubStyle  string            `json:"displaySubStyle"`
}
