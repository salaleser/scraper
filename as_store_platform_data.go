package scraper

// StorePlatformData contains results.
type StorePlatformData struct {
	Results         map[string]Result `json:"results"`
	Version         int               `json:"version"`
	IsAuthenticated bool              `json:"isAuthenticated"`
	Meta            Meta              `json:"meta"`
}
