package scraper

type Metrics struct {
	Config struct{} `json:"config"`
	Fields struct {
		SearchTerm string `json:"searchTerm"`
	} `json:"fields"`
}
